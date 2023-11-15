export APP_ENV=local
export APP_PORT=5000
export MYSQL_USER=root
export MYSQL_PASSWORD=123456
export MYSQL_DATABASE=auth_admin
export MYSQL_PROTOCOL=tcp(127.0.0.1:33061)
export DEV_API_KEY=ema_aria_inc
export ADMIN_JWT_SECRET_KEY=aria_inc_private

run dev:
	go run main.go