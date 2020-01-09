import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    themes: [
      {
        name: "Theme-name1",
        parentTheme: undefined
      },
      {
        name: "Theme-name2",
        parentTheme: "Theme-name1"
      }
    ],
    messages: [
      {
        theme: "Theme-name1",
        authorName: "Meromen",
        text: "asdjkaskjvnkajsdj",
        date: "10.05 19:30"
      },
      {
        theme: "Theme-name1",
        authorName: "Meromen",
        text: "asdjkaskjvnkajsdj",
        date: "10.05 19:30"
      },
      {
        theme: "Theme-name2",
        authorName: "Meromen",
        text: "asdjkaskjvnkajsdj",
        date: "10.05 19:30"
      },
      {
        theme: "Theme-name2",
        authorName: "Meromen",
        text: "asdjkaskjvnkajsdj",
        date: "10.05 19:30"
      }
    ]
  },
  mutations: {},
  actions: {
    GET_THEMES: (contex, payload) => {
      console.log("asd");
      return contex.state.themes.filter(item => {
        if (item.name == payload) return item;
      });
    },
    GET_MESSAGES: (context, payload) => {
      return context.state.messages.filter(item => {
        if (item.theme == payload) return item;
      });
    }
  },
  modules: {}
});
