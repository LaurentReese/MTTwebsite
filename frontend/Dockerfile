# Also useful : https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes-fr

# Dockerfile for VueJS
# Choose the Image which has Node installed already
FROM node:lts-alpine

# installe un simple serveur http pour servir un contenu statique
RUN npm install -g http-server

# définit le dossier 'app' comme dossier de travail
WORKDIR /app

# copie 'package.json' et 'package-lock.json' (si disponible)
COPY package*.json ./

# installe les dépendances du projet
RUN npm install

# copie les fichiers et dossiers du projet dans le dossier de travail (par exemple : le dossier 'app')
COPY . .

# construit l'app pour la production en la minifiant
RUN npm run build

EXPOSE 8080
CMD [ "http-server", "dist" ]

# Now let's build the Docker image of our Vue.js app:
# ==> sudo docker build . -t mtt_frontend
# OLD ==> docker build -t vuejs-cookbook/dockerize-vuejs-app .
# Finally, let's run our Vue.js app in a Docker container:
# ==> docker run -p 8080:8080 -it mtt_frontend
# OLD ==> docker run -it -p 8080:8080 --rm --name dockerize-vuejs-app-1 vuejs-cookbook/dockerize-vuejs-app
# We should be able to access our Vue.js app on localhost:8080.
# delete a docker image
# ==> sudo docker rmi mtt_backend
# purge
# ==> sudo docker images purge
