/// <reference types="vite/client" />
/// <reference types="vue/ref-macros" />
/// <reference types="naive-ui/volar" />
/// <reference types="@types/qrcode" />

interface ImportMetaEnv {
  readonly VITE_HOST: string;
  readonly VITE_ENABLE_ANOUNCEMENT: string;
  readonly VITE_ENABLE_WALLET: string;
  readonly VITE_ALLOW_USER_REGISTER: string;
  readonly VITE_ALLOW_PHONE_BIND: string;
  readonly VITE_ALLOW_ACTIVATION: string;
  readonly VITE_ALLOW_TWEET_ATTACHMENT: string;
  readonly VITE_ALLOW_TWEET_ATTACHMENT_PRICE: string;
  readonly VITE_ALLOW_TWEET_VIDEO: string;
  readonly VITE_ALLOW_TWEET_LABEL: string;
  readonly VITE_ALLOW_TWEET_VISIBILITY: string;
  readonly VITE_DEFAULT_MSG_LOOP_INTERVAL: number;
  readonly VITE_DEFAULT_TWEET_VISIBILITY: string;
  readonly VITE_DEFAULT_TWEET_IMAGE_404: string;
  readonly VITE_DEFAULT_TWEET_MAX_LENGTH: number;
  readonly VITE_DEFAULT_COMMENT_MAX_LENGTH: number;
  readonly VITE_DEFAULT_REPLY_MAX_LENGTH: number;
  readonly VITE_TWEET_IMAGE_THUMBNAIL: string;
  readonly VITE_RIGHT_FOLLOW_TOPIC_MAX_SIZE: number;
  readonly VITE_RIGHT_HOT_TOPIC_MAX_SIZE: number;
  readonly VITE_COPYRIGHT_TOP: string;
  readonly VITE_COPYRIGHT_LEFT: string;
  readonly VITE_COPYRIGHT_LEFT_LINK: string;
  readonly VITE_COPYRIGHT_RIGHT: string;
  readonly VITE_COPYRIGHT_RIGHT_LINK: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
