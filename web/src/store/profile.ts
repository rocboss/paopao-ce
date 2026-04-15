import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useStoreProfile = defineStore('profile', () => {
  const profile = ref({
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
    copyrightTop: '2026 PaoPao',
    copyrightLeft: '',
    copyrightLeftLink: '',
    copyrightRight: 'Github',
    copyrightRightLink: 'https://github.com/rocboss/paopao-ce',
  });

  function loadDefaultSiteProfile() {
    profile.value.useFriendship =
      import.meta.env.VITE_USE_FRIENDSHIP.toLowerCase() === 'true';

    profile.value.enableTrendsBar =
      import.meta.env.VITE_ENABLE_TRENDS_BAR.toLowerCase() === 'true';

    profile.value.enableWallet =
      import.meta.env.VITE_ENABLE_WALLET.toLowerCase() === 'true';

    profile.value.allowTweetAttachment =
      import.meta.env.VITE_ALLOW_TWEET_ATTACHMENT.toLowerCase() === 'true';

    profile.value.allowTweetAttachmentPrice =
      import.meta.env.VITE_ALLOW_TWEET_ATTACHMENT_PRICE.toLowerCase() ===
      'true';

    profile.value.allowTweetVideo =
      import.meta.env.VITE_ALLOW_TWEET_VIDEO.toLowerCase() === 'true';

    profile.value.allowUserRegister =
      import.meta.env.VITE_ALLOW_USER_REGISTER.toLowerCase() === 'true';

    profile.value.allowPhoneBind =
      import.meta.env.VITE_ALLOW_PHONE_BIND.toLowerCase() === 'true';

    profile.value.defaultTweetMaxLength = Number(
      import.meta.env.VITE_DEFAULT_TWEET_MAX_LENGTH,
    );

    profile.value.tweetWebEllipsisSize = Number(
      import.meta.env.VITE_TWEET_WEB_ELLIPSIS_SIZE,
    );

    profile.value.tweetMobileEllipsisSize = Number(
      import.meta.env.VITE_TWEET_MOBILE_ELLIPSIS_SIZE,
    );

    profile.value.defaultTweetVisibility =
      import.meta.env.VITE_DEFAULT_TWEET_VISIBILITY.toLowerCase();

    profile.value.defaultMsgLoopInterval = Number(
      import.meta.env.VITE_DEFAULT_MSG_LOOP_INTERVAL,
    );

    profile.value.copyrightTop = import.meta.env.VITE_COPYRIGHT_TOP;

    profile.value.copyrightLeft = import.meta.env.VITE_COPYRIGHT_LEFT;

    profile.value.copyrightLeftLink = import.meta.env.VITE_COPYRIGHT_LEFT_LINK;

    profile.value.copyrightRight = import.meta.env.VITE_COPYRIGHT_RIGHT;
    profile.value.copyrightRightLink =
      import.meta.env.VITE_COPYRIGHT_RIGHT_LINK;
  }

  function updateSiteProfile(data: Record<string, any>) {
    profile.value.useFriendship =
      data.use_friendship ?? profile.value.useFriendship;

    profile.value.enableTrendsBar =
      data.enable_trends_bar ?? profile.value.enableTrendsBar;

    profile.value.enableWallet =
      data.enable_wallet ?? profile.value.enableWallet;

    profile.value.allowTweetAttachment =
      data.allow_tweet_attachment ?? profile.value.allowTweetAttachment;

    profile.value.allowTweetAttachmentPrice =
      data.allow_tweet_attachment_price ??
      profile.value.allowTweetAttachmentPrice;

    profile.value.allowTweetVideo =
      data.allow_tweet_video ?? profile.value.allowTweetVideo;

    profile.value.allowUserRegister =
      data.allow_user_register ?? profile.value.allowUserRegister;

    profile.value.allowPhoneBind =
      data.allow_phone_bind ?? profile.value.allowPhoneBind;

    profile.value.defaultTweetMaxLength =
      data.default_tweet_max_length ?? profile.value.defaultTweetMaxLength;

    profile.value.tweetWebEllipsisSize =
      data.tweet_web_ellipsis_size ?? profile.value.tweetWebEllipsisSize;

    profile.value.tweetMobileEllipsisSize =
      data.tweet_mobile_ellipsis_size ?? profile.value.tweetMobileEllipsisSize;

    profile.value.defaultTweetVisibility =
      data.default_tweet_visibility ?? profile.value.defaultTweetVisibility;

    profile.value.defaultMsgLoopInterval =
      data.default_msg_loop_interval ?? profile.value.defaultMsgLoopInterval;

    profile.value.copyrightTop =
      data.copyright_top ?? profile.value.copyrightTop;
    profile.value.copyrightLeft =
      data.copyright_left ?? profile.value.copyrightLeft;

    profile.value.copyrightLeftLink =
      data.copyright_left_link ?? profile.value.copyrightLeftLink;

    profile.value.copyrightRight =
      data.copyright_right ?? profile.value.copyrightRight;
    profile.value.copyrightRightLink =
      data.copyright_right_link ?? profile.value.copyrightRightLink;
  }

  return {
    profile,
    loadDefaultSiteProfile,
    updateSiteProfile,
  };
});
