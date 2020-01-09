<template>
  <div class="home">
    <forum-list :themeName="'Theme-name1'"></forum-list>
  </div>
</template>

<script>
import ForumList from "@/components/ForumList";

export default {
  name: "home",
  components: {
    ForumList
  }
};
</script>

<style lang="scss">
.home {
  height: calc(100% - 80px);
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  grid-template-rows: repeat(12, calc(100% / 12));
  grid-template-areas:
    ".  .  .  .  .  .  .  .  .  .  .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  fl fl fl fl fl fl fl fl .  ."
    ".  .  .  .  .  .  .  .  .  .  .  .";
}
</style>
