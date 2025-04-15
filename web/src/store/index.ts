import { createStore } from 'vuex';

export default createStore({
  state: {
    refresh: Date.now(),
    refreshTopicFollow: Date.now(),
    theme: localStorage.getItem('PAOPAO_THEME'),
    collapsedLeft: document.body.clientWidth <= 821,
    collapsedRight: document.body.clientWidth <= 821,
    drawerModelShow: document.body.clientWidth <= 821,
    desktopModelShow: document.body.clientWidth > 821,
    authModalShow: false,
    authModelTab: 'signin',
    unreadMsgCount: 0,
    userLogined: false,
    userInfo: {
      id: 0,
      username: '',
      nickname: '',
      created_on: 0,
      follows: 0,
      followings: 0,
      tweets_count: 0,
      is_admin: false,
    },
    profile: {
      useFriendship: true,
      enableTrendsBar: true,
      enableWallet: false,
      allowTweetAttachment: true,
      allowTweetAttachmentPrice: true,
      allowTweetVideo: true,
      allowUserRegister: true,
      allowPhoneBind: true,
      defaultTweetMaxLength: 2000,
      tweetWebEllipsisSize: 400,
      tweetMobileEllipsisSize: 300,
      defaultTweetVisibility: 'friend',
      defaultMsgLoopInterval: 5000,
      copyrightTop: '2023 paopao.info',
      copyrightLeft: "Roc's Me",
      copyrightLeftLink: '',
      copyrightRight: '泡泡(PaoPao)开源社区',
      copyrightRightLink: 'https://www.paopao.info',
    },
  },
  mutations: {
    refresh(state, refresh) {
      state.refresh = refresh || Date.now();
    },
    refreshTopicFollow(state) {
      state.refreshTopicFollow = Date.now();
    },
    updateUnreadMsgCount(state, count) {
      state.unreadMsgCount = count;
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
    loadDefaultSiteProfile(state) {
      state.profile.useFriendship =
        import.meta.env.VITE_USE_FRIENDSHIP.toLowerCase() === 'true';

      state.profile.enableTrendsBar =
        import.meta.env.VITE_ENABLE_TRENDS_BAR.toLowerCase() === 'true';

      state.profile.enableWallet =
        import.meta.env.VITE_ENABLE_WALLET.toLocaleLowerCase() === 'true';

      state.profile.allowTweetAttachment =
        import.meta.env.VITE_ALLOW_TWEET_ATTACHMENT.toLowerCase() === 'true';

      state.profile.allowTweetAttachmentPrice =
        import.meta.env.VITE_ALLOW_TWEET_ATTACHMENT_PRICE.toLowerCase() ===
        'true';

      state.profile.allowTweetVideo =
        import.meta.env.VITE_ALLOW_TWEET_VIDEO.toLowerCase() === 'true';

      state.profile.allowUserRegister =
        import.meta.env.VITE_ALLOW_USER_REGISTER.toLowerCase() === 'true';

      state.profile.allowPhoneBind =
        import.meta.env.VITE_ALLOW_PHONE_BIND.toLowerCase() === 'true';

      state.profile.defaultTweetMaxLength = Number(
        import.meta.env.VITE_DEFAULT_TWEET_MAX_LENGTH,
      );

      state.profile.tweetWebEllipsisSize = Number(
        import.meta.env.VITE_TWEET_WEB_ELLIPSIS_SIZE,
      );

      state.profile.tweetMobileEllipsisSize = Number(
        import.meta.env.VITE_TWEET_MOBILE_ELLIPSIS_SIZE,
      );

      state.profile.defaultTweetVisibility =
        import.meta.env.VITE_DEFAULT_TWEET_VISIBILITY.toLowerCase();

      state.profile.defaultMsgLoopInterval = Number(
        import.meta.env.VITE_DEFAULT_MSG_LOOP_INTERVAL,
      );

      state.profile.copyrightTop = import.meta.env.VITE_COPYRIGHT_TOP;

      state.profile.copyrightLeft = import.meta.env.VITE_COPYRIGHT_LEFT;

      state.profile.copyrightLeftLink =
        import.meta.env.VITE_COPYRIGHT_LEFT_LINK;

      state.profile.copyrightRight = import.meta.env.VITE_COPYRIGHT_RIGHT;
      state.profile.copyrightRightLink =
        import.meta.env.VITE_COPYRIGHT_RIGHT_LINK;
    },
    updateSiteProfile(state, data) {
      const p = state.profile;
      state.profile.useFriendship = data.use_friendship ?? p.useFriendship;

      state.profile.enableTrendsBar =
        data.enable_trends_bar ?? p.enableTrendsBar;

      state.profile.enableWallet = data.enable_wallet ?? p.enableWallet;

      state.profile.allowTweetAttachment =
        data.allow_tweet_attachment ?? p.allowTweetAttachment;

      state.profile.allowTweetAttachmentPrice =
        data.allow_tweet_attachment_price ?? p.allowTweetAttachmentPrice;

      state.profile.allowTweetVideo =
        data.allow_tweet_video ?? p.allowTweetVideo;

      state.profile.allowUserRegister =
        data.allow_user_register ?? p.allowUserRegister;

      state.profile.allowPhoneBind = data.allow_phone_bind ?? p.allowPhoneBind;

      state.profile.defaultTweetMaxLength =
        data.default_tweet_max_length ?? p.defaultTweetMaxLength;

      state.profile.tweetWebEllipsisSize =
        data.tweet_web_ellipsis_size ?? p.tweetWebEllipsisSize;

      state.profile.tweetMobileEllipsisSize =
        data.tweet_mobile_ellipsis_size ?? p.tweetMobileEllipsisSize;

      state.profile.defaultTweetVisibility =
        data.default_tweet_visibility ?? p.defaultTweetVisibility;

      state.profile.defaultMsgLoopInterval =
        data.default_msg_loop_interval ?? p.defaultMsgLoopInterval;

      state.profile.copyrightTop = data.copyright_top ?? p.copyrightTop;
      state.profile.copyrightLeft = data.copyright_left ?? p.copyrightLeft;

      state.profile.copyrightLeftLink =
        data.copyright_left_link ?? p.copyrightLeftLink;

      state.profile.copyrightRight = data.copyright_right ?? p.copyrightRight;
      state.profile.copyrightRightLink =
        data.copyright_right_link ?? p.copyrightRightLink;
    },
    userLogout(state) {
      localStorage.removeItem('PAOPAO_TOKEN');
      state.userInfo = {
        id: 0,
        nickname: '',
        username: '',
        created_on: 0,
        follows: 0,
        followings: 0,
        tweets_count: 0,
        is_admin: false,
      };
      state.userLogined = false;
    },
  },
  actions: {},
  modules: {},
});
