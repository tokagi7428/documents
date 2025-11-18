FROM node

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

RUN npm install

EXPOSE 4200

CMD [ "ng", "serve" ]