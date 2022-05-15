variable "gcp_zone" {
  type = string
}

job "client_proxy" {
  datacenters = [var.gcp_zone]

  group "nginx" {
    count = 1

    network {
      port "session" {
        static = 3001
      }
    }

    service {
      name = "nginx"
      port = "session"
    }

    task "nginx" {
      driver = "docker"

      config {
        image = "nginx"

        ports = ["session"]

        volumes = [
          "local:/etc/nginx/conf.d",
        ]
      }

      template {
        data = <<EOF
upstream backend {
{{ range service "session-proxy" }}
  server {{ .Address }}:{{ .Port }};
{{ else }}server 127.0.0.1:65535; # force a 502
{{ end }}
}

server {
  listen 3001;
  proxy_set_header Host $host;
  proxy_set_header X-Real-IP $remote_addr;

  location /__health {
      access_log off;
      add_header 'Content-Type' 'text/plain';
      return 200 "healthy\n";
  }

  location / {
    proxy_pass http://backend;
  }
}
EOF

        destination   = "local/load-balancer.conf"
        change_mode   = "signal"
        change_signal = "SIGHUP"
      }
    }
  }
}
