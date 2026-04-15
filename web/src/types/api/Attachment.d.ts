declare namespace Api {
  namespace Attachment {
    interface Api {
      get: {
        /** 获取附件 */
        _self: (
          params: NetParams.UserGetAttachment,
        ) => Promise<NetReq.UserGetAttachment>;
        /** 获取附件 */
        precheck: (
          params: NetParams.UserPrecheckAttachment,
        ) => Promise<NetReq.UserPrecheckAttachment>;
      };
    }

    namespace NetParams {
      interface UserGetAttachment {
        id: number;
      }

      interface UserPrecheckAttachment {
        id: number;
      }
    }

    namespace NetReq {
      interface UserGetAttachment {
        signed_url: string;
      }

      interface UserPrecheckAttachment {
        paid: number;
      }
    }
  }
}
