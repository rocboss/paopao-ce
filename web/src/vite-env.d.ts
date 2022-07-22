/// <reference types="vite/client" />
/// <reference types="vue/ref-macros" />
/// <reference types="naive-ui/volar" />
/// <reference types="@types/qrcode" />

interface ImportMetaEnv {
    readonly VITE_HOST: string
    readonly VITE_ENABLE_WALLET: string
    readonly VITE_ALLOW_TWEET_ATTACHMENT: string
    readonly VITE_ALLOW_TWEET_ATTACHMENT_PRICE: string
    readonly VITE_ALLOW_TWEET_VIDEO: string
    readonly VITE_ALLOW_TWEET_VISIBILITY: string
    readonly VITE_ALLOW_PHONE_BIND: string
    readonly VITE_DEFAULT_TWEET_VISIBILITY: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}