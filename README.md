# TheBlock
Make human connections by putting your personality first!

Unlike other social networking apps, we believe in a swipe-less world. In real life, we often make friends with people we harshly judge first or may even have little to nothing in common with. By swiping left on those people, we eliminate those special connections that only happen by chance. As each new generation grows up with newer technologies, more and more people are meeting online and yet there's a stigma that we don't meet "organically". The internet is a tool to help bring us together, not something we should keep a secret in our closet. The Block wants to recreate those organic connections by allowing you to do you best--by putting YOURSELF out there first. Not a curated profile of perfect prompt answers and hand-selected photos.

You never know who you are going to meet around the block :)

# Local Development
To start local development please run:
```make start```
This will start react and golang servers and also run migrations on a docker database container.

## Needed Dependencies
Please ensure you have the following installed before starting local development:

### for local env
- docker
- golang-migrate
- react
- golang
- sql

### react packages
Install using `npm install -save [package name]` in ./react-app/the-block
- styled-components
- react-router-dom
