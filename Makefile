migrate:
	goose -dir migrations postgres "postgresql://pguser:pgpassword@127.0.0.1:7222/postgres?sslmode=disable" up

create_posts:
	goose -dir migrations create posts sql 

