migrate:
	migrate -path db/migrations -database "postgres://default:secret@127.0.0.1:5432/final-project?sslmode=disable" up 

rollback:
	migrate -database "postgres://default:secret@127.0.0.1:5432/final-project?sslmode=disable" -path db/migrations  drop -f

test: 
	go env