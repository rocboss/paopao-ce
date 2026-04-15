interface Api {
  v1: {
    auth: Api.Auth.Api;
    user: Api.User.Api;
    friend: Api.Friend.Api;
    captcha: Api.Captcha.Api;
    attachment: Api.Attachment.Api;

    suggest: Api.Suggest.Api;
    admin: Api.Admin.Api;
  };
}
