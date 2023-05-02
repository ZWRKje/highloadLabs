FROM node:18

WORKDIR /usr/src/auth

COPY package*.json ./
RUN npm install
COPY . .
# EXPOSE 50051

# CMD [ "node", "servise.js" ]