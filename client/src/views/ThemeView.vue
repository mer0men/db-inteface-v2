<template>
  <div class="theme-view">
    <forum-list>
      <list-item :name="'Взрыв на станции'" :authorName="'Meromen'" :lastMessage="'07.08.1956'"></list-item>
      <list-item :name="'Взрыв на станции'" :authorName="'Meromen'" :lastMessage="'07.08.1956'"></list-item>
      <list-item :name="'Взрыв на станции'" :authorName="'Meromen'" :lastMessage="'07.08.1956'"></list-item>
      <list-item :name="'Взрыв на станции'" :authorName="'Meromen'" :lastMessage="'07.08.1956'"></list-item>
      <list-item :name="'Взрыв на станции'" :authorName="'Meromen'" :lastMessage="'07.08.1956'"></list-item>
    </forum-list>
  </div>
</template>

<script>
import ForumList from "@/components/ForumList";

export default {
  name: "theme-view",
  components: {
    ForumList
  }
};
</script>

<style lang="scss" >
.theme-view {
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
    ".  .  .  .  .  .  .  .  .  .  .  ."
    ".  .  ml ml ml ml ml ml ml ml .  ."
    ".  .  ml ml ml ml ml ml ml ml .  ."
    ".  .  ml ml ml ml ml ml ml ml .  ."
    ".  .  ml ml ml ml ml ml ml ml .  ."
    ".  .  .  .  .  .  .  .  .  .  .  .";
}
</style>