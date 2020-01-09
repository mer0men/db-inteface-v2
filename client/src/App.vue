<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Главная</router-link>
      <router-link to="/newtheme">Создать тему</router-link>
    </div>
    <router-view />
  </div>
</template>

<style lang="scss">
@import "Variables.scss";

body,
html {
  margin: 0;
  padding: 0;
  height: 100vh;
  position: relative;
}

#app {
  font-family: "Exo 2", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: $text-color;
  background: $background-color;
  height: 100%;
  position: relative;
}

#nav {
  padding: 0 32px;
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  background: $text-color;
  a {
    font-weight: bold;
    color: $background-color;
    margin: 0 16px;

    &.router-link-exact-active {
      color: $active-color;
    }
  }
}
</style>
