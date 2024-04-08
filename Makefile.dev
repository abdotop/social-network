# Makefile

install:
	# Installer les dépendances du back-end
	cd backend && go get

	# Installer les dépendances du front-end
	cd frontend && npm install

start-backend:
	# Démarrer le serveur back-end
	cd backend && ./watch-run.sh -p "*.go"

start-frontend:
	# Démarrer le serveur front-end
	cd frontend && npm run dev

start:
	# Démarrer les deux serveurs en même temps
	make start-backend & make start-frontend