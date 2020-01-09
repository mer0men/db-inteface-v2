<template>
  <div class="forum-list">
    <list-item
      v-for="(item, index) in themes"
      :key="index"
      :name="item.name"
      :authorName="'Meromen'"
      :lastMessage="'07.08.1956'"
    ></list-item>
  </div>
</template>

<script>
import ListItem from "@/components/ListItem";

export default {
  props: {
    themeName: {
      required: true
    }
  },
  components: {
    ListItem
  },
  data() {
    return {
      themes: []
    };
  },
  mounted() {
    this.$store.dispatch("GET_THEMES", this.themeName).then(data => {
      this.themes = data;
    });
  }
};
</script>

<style lang="scss">
.forum-list {
  grid-area: fl;
  position: relative;
  overflow-y: scroll;
}
</style>