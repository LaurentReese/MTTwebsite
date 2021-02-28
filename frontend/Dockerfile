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
# ==> docker build . -t mtt-frontend
# see the results :
# ==> docker images
# Finally, let's run our Vue.js app in a Docker container:
# ==> docker run -p 8080:8080 -it mtt-frontend
# delete a docker image (add --force to force)
# ==> docker rmi mtt-frontend
# purge
# ==> docker images purge
#
# To build a distant image (will go to Docker Hub, 32681733 is my Docker Hub ID)
# ==> docker build . -t 32681733/mtt-frontend
# To test locally the distant Docker image
# ==> docker run -it -p 8080:8080 -m 512m --memory-swap=512m 32681733/mtt-frontend
# To push the image to Docker Hub (32681733 is my Docker Hub ID)
# ==> docker push 32681733/mtt-frontend




