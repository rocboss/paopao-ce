import { useDialog } from "naive-ui";
import { Api } from "./request";

const dialog = useDialog();

export default class UserAction {

    static followAction(userId: number, userName: string, isFollowing: boolean) {
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

}