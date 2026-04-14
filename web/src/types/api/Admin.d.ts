declare namespace Api {

    namespace Admin {

        interface Api {
            post: {
                user: {
                    /** 管理·用户禁言/解禁 */
                    status: (params: NetParams.UserStatusReq) => Promise<NetReq.UserChangeStatus>;
                },
                site: {
                    /** 管理·更新系统配置 */
                    profile: (params: NetParams.SiteProfileReq) => Promise<NetReq.SiteProfileResp>;
                },
                settings: {
                    /** 管理·保存通用配置 */
                    save: (params: NetParams.SettingsSaveReq) => Promise<NetReq.SettingsSaveResp>;
                }
            },
            get: {
                site: {
                    /** 获取站点状态信息 */
                    status: () => Promise<NetReq.SiteInfoResp>;
                    /** 管理·获取系统配置 */
                    profile: () => Promise<NetReq.SiteProfileResp>;
                },
                settings: {
                    /** 管理·获取通用配置 schema */
                    schema: () => Promise<NetReq.SettingsSchemaResp>;
                    /** 管理·获取通用配置当前值 */
                    values: () => Promise<NetReq.SettingsValuesResp>;
                }
            }
        }

        namespace NetParams {

            interface UserStatusReq {
                id: number;
                status: number;
            }

            interface SiteProfileReq {
                use_friendship: boolean;
                enable_trends_bar: boolean;
                enable_wallet: boolean;
                allow_tweet_attachment: boolean;
                allow_tweet_attachment_price: boolean;
                allow_tweet_video: boolean;
                default_tweet_max_length: number;
                tweet_web_ellipsis_size: number;
                tweet_mobile_ellipsis_size: number;
                default_tweet_visibility: 'public' | 'following' | 'friend' | 'private';
                default_msg_loop_interval: number;
                copyright_top: string;
                copyright_left: string;
                copyright_left_link: string;
                copyright_right: string;
                copyright_right_link: string;
            }

            interface SettingValueInput {
                key: string;
                value: string | number | boolean;
            }

            interface SettingsSaveReq {
                items: SettingValueInput[];
            }

        }

        namespace NetReq {
            interface UserChangeStatus {}

            interface SettingOption {
                label: string;
                value: string | number | boolean;
            }

            interface SettingsSchemaItem {
                key: string;
                group: string;
                section: string;
                type: 'bool' | 'int' | 'float' | 'string';
                label: string;
                description: string;
                apply_mode: 'live' | 'restart_required' | 'bootstrap_only';
                secret: boolean;
                readonly: boolean;
                active: boolean;
                bootstrap_value?: string | number | boolean;
                bootstrap_configured?: boolean;
                options?: SettingOption[];
            }

            interface SettingsSchemaResp {
                items: SettingsSchemaItem[];
            }

            interface SettingsValueItem {
                key: string;
                value?: string | number | boolean;
                effective_value?: string | number | boolean;
                source: 'bootstrap' | 'override' | string;
                pending_restart: boolean;
                configured?: boolean;
                active: boolean;
            }

            interface SettingsValuesResp {
                items: SettingsValueItem[];
                has_pending_restart: boolean;
            }

            interface SettingsSaveResp {
                items: SettingsValueItem[];
                updated_keys: string[];
                has_pending_restart: boolean;
            }

            interface SiteInfoResp {
                register_user_count: number;
                online_user_count: number;
                history_max_online: number;
                server_up_time: number;
            }

            interface SiteProfileResp {
                use_friendship: boolean;
                enable_trends_bar: boolean;
                enable_wallet: boolean;
                allow_tweet_attachment: boolean;
                allow_tweet_attachment_price: boolean;
                allow_tweet_video: boolean;
                allow_user_register: boolean;
                allow_phone_bind: boolean;
                default_tweet_max_length: number;
                tweet_web_ellipsis_size: number;
                tweet_mobile_ellipsis_size: number;
                default_tweet_visibility: 'public' | 'following' | 'friend' | 'private';
                default_msg_loop_interval: number;
                copyright_top: string;
                copyright_left: string;
                copyright_left_link: string;
                copyright_right: string;
                copyright_right_link: string;
                readonly_fields: string[];
            }
        }

    }

}
