import { stat } from "fs";
import { createStore } from "vuex";

export default createStore({
  state: {
    refresh: Date.now(),
    refreshTopicFollow: Date.now(),
    theme: localStorage.getItem("PAOPAO_THEME"),
    collapsedLeft: document.body.clientWidth <= 821,
    collapsedRight: document.body.clientWidth <= 821,
    drawerModelShow: document.body.clientWidth <= 821,
    desktopModelShow: document.body.clientWidth > 821,
    authModalShow: false,
    authModelTab: "signin",
    userLogined: false,
    userInfo: {
      id: 0,
      username: "",
      nickname: "",
      created_on: 0,
      follows: 0,
      followings: 0,
    },
  },
  mutations: {
    refresh(state, refresh) {
      state.refresh = refresh || Date.now();
    },
    refreshTopicFollow(state) {
      state.refreshTopicFollow = Date.now();
    },
    triggerTheme(state, theme) {
      state.theme = theme;
    },
    triggerAuth(state, status) {
      state.authModalShow = status;
    },
    triggerAuthKey(state, key) {
      state.authModelTab = key;
    },
    triggerCollapsedLeft(state, status) {
      state.collapsedLeft = status;
      state.drawerModelShow = status;
      state.desktopModelShow = !status;
    },
    triggerCollapsedRight(state, status) {
      state.collapsedRight = status;
    },
    updateUserinfo(state, data) {
      state.userInfo = data;
      if (state.userInfo.id > 0) {
        state.userLogined = true;
      }
    },
    userLogout(state) {
      localStorage.removeItem("PAOPAO_TOKEN");
      state.userInfo = {
        id: 0,
        nickname: "",
        username: "",
        created_on: 0,
        follows: 0,
        followings: 0,
      };
      state.userLogined = false;
    },
  },
  actions: {},
  modules: {},
});
