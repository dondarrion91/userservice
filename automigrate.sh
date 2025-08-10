atlas migrate diff --env gorm
atlas migrate apply \
  --dir "file://migrations" \
  --url "mysql://root:root@127.0.0.1:3306/user_service"