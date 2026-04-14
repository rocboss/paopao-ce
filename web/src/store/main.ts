import { defineStore } from 'pinia';
import { reactive, ref } from 'vue';

export const useStoreMain = defineStore('main', () => {
  const refresh = ref(Date.now());
  const refreshTopicFollow = ref(Date.now());
  const theme = ref(localStorage.getItem('PAOPAO_THEME'));
  const collapsedLeft = ref(document.body.clientWidth <= 821);
  const collapsedRight = ref(document.body.clientWidth <= 821);
  const drawerModelShow = ref(document.body.clientWidth <= 821);
  const desktopModelShow = ref(document.body.clientWidth > 821);
  const authModalShow = ref(false);
  const authModelTab = ref('signin');
  const unreadMsgCount = ref(0);

  function doRefresh(val?: number) {
    refresh.value = val || Date.now();
  }

  function doRefreshTopicFollow() {
    refreshTopicFollow.value = Date.now();
  }

  function updateUnreadMsgCount(count: number) {
    unreadMsgCount.value = count;
  }

  function triggerTheme(t: string) {
    theme.value = t;
  }

  function triggerAuth(status: boolean) {
    authModalShow.value = status;
  }

  function triggerAuthKey(key: string) {
    authModelTab.value = key;
  }

  function triggerCollapsedLeft(status: boolean) {
    collapsedLeft.value = status;
    drawerModelShow.value = status;
    desktopModelShow.value = !status;
  }

  function triggerCollapsedRight(status: boolean) {
    collapsedRight.value = status;
  }

  return {
    refresh,
    refreshTopicFollow,
    theme,
    collapsedLeft,
    collapsedRight,
    drawerModelShow,
    desktopModelShow,
    authModalShow,
    authModelTab,
    unreadMsgCount,
    doRefresh,
    doRefreshTopicFollow,
    updateUnreadMsgCount,
    triggerTheme,
    triggerAuth,
    triggerAuthKey,
    triggerCollapsedLeft,
    triggerCollapsedRight,
  };
});
