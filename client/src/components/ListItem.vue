<template>
  <div class="forum-list__item">
    <span class="name">{{ name }}</span>
    <span class="author-label">Автор:</span>
    <span class="author">{{ authorName }}</span>
    <span class="last-message-label">Последние сообщение:</span>
    <span class="last-message">{{ lastMessage }}</span>
  </div>
</template>

<script>
export default {
  props: {
    name: {
      required: true
    },
    authorName: {
      required: true
    },
    lastMessage: {
      required: true
    }
  }
};
</script>

<style lang="scss" scoped>
@import "../Variables.scss";

.forum-list__item {
  padding: 6px;
  display: grid;
  background: $text-color;
  color: $background-color;
  border: 2px solid $action-color;
  border-radius: 8px;
  margin: 8px 0;
  grid-template-columns: repeat(12, 1fr);
  grid-template-rows: 32px 32px;
  transition: 0.4s;
  grid-template-areas:
    "nm  nm  nm  nm  nm  nm  nm  nm  al  al  ml  ml"
    "nm  nm  nm  nm  nm  nm  nm  nm  an  an  lm  lm";

  cursor: pointer;

  &:hover {
    border-color: $active-color;
  }

  .name {
    grid-area: nm;
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .author-label {
    grid-area: al;
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .author {
    grid-area: an;
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .last-message-label {
    grid-area: ml;
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .last-message {
    grid-area: lm;
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }
}
</style>