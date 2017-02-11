# Go + Javascript bootstrap template.

An opinionated bootstrap project template to create simple SPA client-server applications with Javascript frontend and Go language backend.

It uses Webpack to create Javascript app bundles and [statik](https://github.com/rakyll/statik) tool to embed resulting application into Go binary. As result you get ready to deploy single self-contained binary that will work without dependencies everywhere.

For example deploy to Heroku in one click:

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/FZambia/go-javascript-template)

And get something [like this](https://go-javascript-template.herokuapp.com/).

## Setup

To use this bootstrap you’ll need to install a few things before you have a working copy of the project.

### 1. Install Go language and statik tool

https://golang.org/

```
go get -u github.com/rakyll/statik
```

### 2. Clone this repo:

Navigate into your workspace directory.

Run:

```git clone https://github.com/FZambia/go-javascript-webpack.git```

### 3. Install node.js and npm:

https://nodejs.org/en/


### 4. Install dependencies:

Navigate to the cloned repo’s directory.

Run:

```npm install```

### 5. Run the development server:

Run:

```npm run dev```

This will run a server so you can run the app in a browser.

Open your browser and enter localhost:3000 into the address bar.

Also this will start a watch process, so you can change the source and the process will recompile and refresh the browser

## Build for deployment:

Run:

```npm run deploy```

This will optimize and minimize the compiled bundle.

## Test Javascript app

```
npm test
```

## Run Go server with builtin web app

```
go run main.go
```

## Serve app that was built with Webpack using Go

```
go run main.go --web_path=app/dist
```

## Update built-in web app in Go

```
npm run deploy
statik -src=app/dist -dest ./
```

## Features

- Webpack for build process
- ES6 support using Babel
- Multiple browser testing using BrowserSync
- WebFont Loader
- ESLINT with JavaScript Standard Style configuration

## Credits

Big thanks to this repo:

https://github.com/lean/phaser-es6-webpack
