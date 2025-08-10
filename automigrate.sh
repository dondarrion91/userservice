atlas migrate diff --env gorm
atlas migrate apply \
  --dir "file://migrations" \
  --url "mysql://root:root@172.17.0.2:3306/user_service"