.PHONY: migration/lint
migration/lint:
	go run -mod=mod ariga.io/atlas/cmd/atlas@v0.7.0 migrate lint --dev-url="sqlite://main.db" --dir="file://ent/migrate/migration" --latest=1

.PHONY: migration/run
migration/run:
	go run -mod=mod ariga.io/atlas/cmd/atlas@v0.7.0 migrate apply --dir="file://ent/migrate/migration" --url="sqlite://main.db?_fk=1"
