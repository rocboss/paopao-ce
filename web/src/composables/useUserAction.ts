import { ref } from 'vue';
import { useDialog } from "naive-ui";
import { Api } from "../utils/request";

export default class UserAction {

    /**
     * 用户关注/取消关注操作
     * @param dialog dialog 实例（从组件中传入）
     * @param userId 目标用户ID
     * @param userName 目标用户名（用于提示）
     * @param isFollowing 当前是否已关注（用于确定操作类型和提示文案）
     */
    static followAction(dialog: ReturnType<typeof useDialog>, userId: number, userName: string, isFollowing: boolean) {
        return new Promise<boolean>((resolve, reject) => {
            dialog.success({
                title: '提示',
                content:
                    '确定' +
                    (isFollowing ? '取消关注 @' : '关注 @') +
                    userName +
                    ' 吗？',
                positiveText: '确定',
                negativeText: '取消',
                onPositiveClick: () => {
                    if (isFollowing) {
                        Api.v1.user.post.unfollow({
                            user_id: userId,
                        })
                        .then((_res) => {
                            window.$message.success('操作成功');
                            resolve(false);
                        })
                        .catch((_err) => {
                            reject(_err);
                        });
                    } else {
                        Api.v1.user.post.follow({
                            user_id: userId,
                        })
                        .then((_res) => {
                            window.$message.success('关注成功');
                            resolve(true);
                        })
                        .catch((_err) => {
                            reject(_err);
                        });
                    }
                },
            });
        });
    }

    /**
     * 用户私信相关操作
     * 保持与原有代码完全一致的行为
     */
    static useWhisper() {
        const showWhisper = ref(false);
        const whisperReceiver = ref<Item.UserInfo>({
            id: 0,
            avatar: '',
            username: '',
            nickname: '',
            is_admin: false,
            is_friend: true,
            is_following: false,
            created_on: 0,
            follows: 0,
            followings: 0,
            status: 1,
        });

        const onSendWhisper = (user: Item.UserInfo) => {
            whisperReceiver.value = user;
            showWhisper.value = true;
        };

        const whisperSuccess = () => {
            showWhisper.value = false;
        };

        return {
            showWhisper,
            whisperReceiver,
            onSendWhisper,
            whisperSuccess,
        };
    }

}
