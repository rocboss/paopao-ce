import{_ as we}from"./main-nav.vue_vue_type_style_index_0_lang-d6d2ed7f.js";import{d as ye,r as d,D as H,j as ke,c as g,L as t,Y as s,K as b,O as _,o as r,a as m,U as u,e as c,M as R,z as be,v as Ce,P as q,Z as B,a4 as Ie,a5 as $e}from"./@vue-f70ab1bd.js";import{u as Pe}from"./vuex-cc1858c6.js";import{L as X,M as Ue,N as Se,O as Re,P as qe,Q as Be,R as Ae,_ as Ne}from"./index-ce5b62d8.js";import{s as ze}from"./@vicons-fc06a0bb.js";import{h as Ke,o as De,e as xe,B as Fe,b as Oe,j as je,S as Le,$ as Me,J as Te,a0 as Ve,a1 as Ee,d as Je}from"./naive-ui-ddb574dd.js";import"./vue-router-29025daf.js";import"./vooks-dfdd6eef.js";import"./evtd-b614532e.js";import"./axios-707ed124.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-804c4158.js";import"./@css-render-66126308.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";const O=A=>(Ie("data-v-a681720e"),A=A(),$e(),A),We={class:"base-line avatar"},Ge={class:"base-line"},Qe=O(()=>m("span",{class:"base-label"},"昵称",-1)),Ye={key:0},Ze={class:"base-line"},He=O(()=>m("span",{class:"base-label"},"用户名",-1)),Xe={key:0},et={key:1},tt=O(()=>m("br",null,null,-1)),at={key:2,class:"phone-bind-wrap"},st={class:"captcha-img-wrap"},nt={class:"captcha-img"},ot=["src"],lt={class:"form-submit-wrap"},rt={key:0},it={key:1},ut=O(()=>m("br",null,null,-1)),dt={key:2,class:"phone-bind-wrap"},ct={class:"captcha-img-wrap"},pt={class:"captcha-img"},_t=["src"],mt={class:"form-submit-wrap"},vt={key:1,class:"phone-bind-wrap"},ft={class:"form-submit-wrap"},gt=ye({__name:"Setting",setup(A){const ee="/v1/attachment",te="Bearer "+localStorage.getItem("PAOPAO_TOKEN"),N=d("public/avatar"),$="true".toLowerCase()==="true",ae="false".toLowerCase()==="true",o=Pe(),P=d(!1),z=d(!1),K=d(!1),E=d(),J=d(),C=d(!1),D=d(!1),U=d(!1),S=d(!1),I=d(60),y=d(!1),k=d(!1),W=d(),G=d(),Q=d(),Y=d(),a=H({id:"",b64s:"",imgCaptcha:"",phone:"",phone_captcha:"",password:"",old_password:"",reenteredPassword:""}),i=H({id:"",b64s:"",imgCaptcha:"",activate_code:""}),se=async n=>{var e,v;return N.value==="public/avatar"&&!["image/png","image/jpg","image/jpeg"].includes((e=n.file.file)==null?void 0:e.type)?(window.$message.warning("头像仅允许 png/jpg 格式"),!1):N.value==="image"&&((v=n.file.file)==null?void 0:v.size)>1048576?(window.$message.warning("头像大小不能超过1MB"),!1):!0},ne=({file:n,event:e})=>{var v;try{let f=JSON.parse((v=e.target)==null?void 0:v.response);f.code===0&&N.value==="public/avatar"&&Ue({avatar:f.data.content}).then(p=>{var x;window.$message.success("头像更新成功"),(x=E.value)==null||x.clear(),o.commit("updateUserinfo",{...o.state.userInfo,avatar:f.data.content})}).catch(p=>{console.log(p)})}catch{window.$message.error("上传失败")}},oe=(n,e)=>!!a.password&&a.password.startsWith(e)&&a.password.length>=e.length,le=(n,e)=>e===a.password,re=()=>{var n;a.reenteredPassword&&((n=Y.value)==null||n.validate({trigger:"password-input"}))},ie=n=>{var e;n.preventDefault(),(e=Q.value)==null||e.validate(v=>{v||(D.value=!0,Se({password:a.password,old_password:a.old_password}).then(f=>{D.value=!1,U.value=!1,window.$message.success("密码重置成功"),o.commit("userLogout"),o.commit("triggerAuth",!0),o.commit("triggerAuthKey","signin")}).catch(f=>{D.value=!1}))})},ue=n=>{var e;n.preventDefault(),(e=W.value)==null||e.validate(v=>{v||(z.value=!0,Re({phone:a.phone,captcha:a.phone_captcha}).then(f=>{z.value=!1,y.value=!1,window.$message.success("绑定成功"),o.commit("updateUserinfo",{...o.state.userInfo,phone:a.phone}),a.id="",a.b64s="",a.imgCaptcha="",a.phone="",a.phone_captcha=""}).catch(f=>{z.value=!1}))})},de=n=>{var e;n.preventDefault(),(e=G.value)==null||e.validate(v=>{if(i.imgCaptcha===""){window.$message.warning("请输入图片验证码");return}P.value=!0,v||(K.value=!0,qe({activate_code:i.activate_code,captcha_id:i.id,imgCaptcha:i.imgCaptcha}).then(f=>{K.value=!1,k.value=!1,window.$message.success("激活成功"),o.commit("updateUserinfo",{...o.state.userInfo,activation:i.activate_code}),i.id="",i.b64s="",i.imgCaptcha="",i.activate_code=""}).catch(f=>{K.value=!1,f.code===20012&&L()}))})},j=()=>{X().then(n=>{a.id=n.id,a.b64s=n.b64s}).catch(n=>{console.log(n)})},L=()=>{X().then(n=>{i.id=n.id,i.b64s=n.b64s}).catch(n=>{console.log(n)})},ce=()=>{Be({nickname:o.state.userInfo.nickname||""}).then(n=>{C.value=!1,window.$message.success("昵称修改成功")}).catch(n=>{C.value=!0})},pe=()=>{if(!(I.value>0&&S.value)){if(a.imgCaptcha===""){window.$message.warning("请输入图片验证码");return}P.value=!0,Ae({phone:a.phone,img_captcha:a.imgCaptcha,img_captcha_id:a.id}).then(n=>{S.value=!0,P.value=!1,window.$message.success("发送成功");let e=setInterval(()=>{I.value--,I.value===0&&(clearInterval(e),I.value=60,S.value=!1)},1e3)}).catch(n=>{P.value=!1,n.code===20012&&j(),console.log(n)})}},_e={phone:[{required:!0,message:"请输入手机号",trigger:["input"],validator:(n,e)=>/^[1]+[3-9]{1}\d{9}$/.test(e)}],phone_captcha:[{required:!0,message:"请输入手机验证码"}]},me={activate_code:[{required:!0,message:"请输入激活码",trigger:["input"],validator:(n,e)=>/\d{6}$/.test(e)}]},ve={password:[{required:!0,message:"请输入新密码"}],old_password:[{required:!0,message:"请输入旧密码"}],reenteredPassword:[{required:!0,message:"请再次输入密码",trigger:["input","blur"]},{validator:oe,message:"两次密码输入不一致",trigger:"input"},{validator:le,message:"两次密码输入不一致",trigger:["blur","password-input"]}]},fe=()=>{C.value=!0,setTimeout(()=>{var n;(n=J.value)==null||n.focus()},30)};return ke(()=>{o.state.userInfo.id===0&&(o.commit("triggerAuth",!0),o.commit("triggerAuthKey","signin")),j(),L()}),(n,e)=>{const v=we,f=De,p=xe,x=Fe,h=Oe,ge=je,F=Ke,Z=Le,w=Me,he=Te,M=Ve,T=Ee,V=Je;return r(),g("div",null,[t(v,{title:"设置",theme:""}),t(F,{title:"基本信息",size:"small",class:"setting-card"},{default:s(()=>[m("div",We,[t(f,{class:"avatar-img",size:80,src:u(o).state.userInfo.avatar},null,8,["src"]),!$||$&&u(o).state.userInfo.phone&&u(o).state.userInfo.phone.length>0?(r(),b(x,{key:0,ref_key:"avatarRef",ref:E,action:ee,headers:{Authorization:te},data:{type:N.value},onBeforeUpload:se,onFinish:ne},{default:s(()=>[t(p,{size:"small"},{default:s(()=>[c("更改头像")]),_:1})]),_:1},8,["headers","data"])):_("",!0)]),m("div",Ge,[Qe,C.value?_("",!0):(r(),g("div",Ye,R(u(o).state.userInfo.nickname),1)),be(t(h,{ref_key:"inputInstRef",ref:J,class:"nickname-input",value:u(o).state.userInfo.nickname,"onUpdate:value":e[0]||(e[0]=l=>u(o).state.userInfo.nickname=l),type:"text",size:"small",placeholder:"请输入昵称",onBlur:ce,maxlength:16},null,8,["value"]),[[Ce,C.value]]),!C.value&&(!$||$&&u(o).state.userInfo.phone&&u(o).state.userInfo.phone.length>0&&u(o).state.userInfo.status==1)?(r(),b(p,{key:1,quaternary:"",round:"",type:"success",size:"small",onClick:fe},{icon:s(()=>[t(ge,null,{default:s(()=>[t(u(ze))]),_:1})]),_:1})):_("",!0)]),m("div",Ze,[He,c(" @"+R(u(o).state.userInfo.username),1)])]),_:1}),$?(r(),b(F,{key:0,title:"手机号",size:"small",class:"setting-card"},{default:s(()=>[u(o).state.userInfo.phone&&u(o).state.userInfo.phone.length>0?(r(),g("div",Xe,[c(R(u(o).state.userInfo.phone)+" ",1),!y.value&&u(o).state.userInfo.status==1?(r(),b(p,{key:0,quaternary:"",round:"",type:"success",onClick:e[1]||(e[1]=l=>y.value=!0)},{default:s(()=>[c(" 换绑手机 ")]),_:1})):_("",!0)])):(r(),g("div",et,[t(Z,{title:"手机绑定提示",type:"warning"},{default:s(()=>[c(" 成功绑定手机后，才能进行换头像、发动态、回复等交互~"),tt,y.value?_("",!0):(r(),g("a",{key:0,class:"hash-link",onClick:e[2]||(e[2]=l=>y.value=!0)}," 立即绑定 "))]),_:1})])),y.value?(r(),g("div",at,[t(V,{ref_key:"phoneFormRef",ref:W,model:a,rules:_e},{default:s(()=>[t(w,{path:"phone",label:"手机号"},{default:s(()=>[t(h,{value:a.phone,"onUpdate:value":e[3]||(e[3]=l=>a.phone=l.trim()),placeholder:"请输入中国大陆手机号",onKeydown:e[4]||(e[4]=q(B(()=>{},["prevent"]),["enter"]))},null,8,["value"])]),_:1}),t(w,{path:"img_captcha",label:"图形验证码"},{default:s(()=>[m("div",st,[t(h,{value:a.imgCaptcha,"onUpdate:value":e[5]||(e[5]=l=>a.imgCaptcha=l),placeholder:"请输入图形验证码后获取验证码"},null,8,["value"]),m("div",nt,[a.b64s?(r(),g("img",{key:0,src:a.b64s,onClick:j},null,8,ot)):_("",!0)])])]),_:1}),t(w,{path:"phone_captcha",label:"短信验证码"},{default:s(()=>[t(he,null,{default:s(()=>[t(h,{value:a.phone_captcha,"onUpdate:value":e[6]||(e[6]=l=>a.phone_captcha=l),placeholder:"请输入收到的短信验证码"},null,8,["value"]),t(p,{type:"primary",ghost:"",disabled:S.value,loading:P.value,onClick:pe},{default:s(()=>[c(R(I.value>0&&S.value?I.value+"s后重新发送":"发送验证码"),1)]),_:1},8,["disabled","loading"])]),_:1})]),_:1}),t(T,{gutter:[0,24]},{default:s(()=>[t(M,{span:24},{default:s(()=>[m("div",lt,[t(p,{quaternary:"",round:"",onClick:e[7]||(e[7]=l=>y.value=!1)},{default:s(()=>[c(" 取消 ")]),_:1}),t(p,{secondary:"",round:"",type:"primary",loading:z.value,onClick:ue},{default:s(()=>[c(" 绑定 ")]),_:1},8,["loading"])])]),_:1})]),_:1})]),_:1},8,["model"])])):_("",!0)]),_:1})):_("",!0),ae?(r(),b(F,{key:1,title:"激活码",size:"small",class:"setting-card"},{default:s(()=>[u(o).state.userInfo.activation&&u(o).state.userInfo.activation.length>0?(r(),g("div",rt,[c(R(u(o).state.userInfo.activation)+" ",1),k.value?_("",!0):(r(),b(p,{key:0,quaternary:"",round:"",type:"success",onClick:e[8]||(e[8]=l=>k.value=!0)},{default:s(()=>[c(" 重新激活 ")]),_:1}))])):(r(),g("div",it,[t(Z,{title:"激活码激活提示",type:"warning"},{default:s(()=>[c(" 成功激活后后，才能发（公开/好友可见）动态、回复~"),ut,k.value?_("",!0):(r(),g("a",{key:0,class:"hash-link",onClick:e[9]||(e[9]=l=>k.value=!0)}," 立即激活 "))]),_:1})])),k.value?(r(),g("div",dt,[t(V,{ref_key:"activateFormRef",ref:G,model:i,rules:me},{default:s(()=>[t(w,{path:"activate_code",label:"激活码"},{default:s(()=>[t(h,{value:i.activate_code,"onUpdate:value":e[10]||(e[10]=l=>i.activate_code=l.trim()),placeholder:"请输入激活码",onKeydown:e[11]||(e[11]=q(B(()=>{},["prevent"]),["enter"]))},null,8,["value"])]),_:1}),t(w,{path:"img_captcha",label:"图形验证码"},{default:s(()=>[m("div",ct,[t(h,{value:i.imgCaptcha,"onUpdate:value":e[12]||(e[12]=l=>i.imgCaptcha=l),placeholder:"请输入图形验证码后获取验证码"},null,8,["value"]),m("div",pt,[i.b64s?(r(),g("img",{key:0,src:i.b64s,onClick:L},null,8,_t)):_("",!0)])])]),_:1}),t(T,{gutter:[0,24]},{default:s(()=>[t(M,{span:24},{default:s(()=>[m("div",mt,[t(p,{quaternary:"",round:"",onClick:e[13]||(e[13]=l=>k.value=!1)},{default:s(()=>[c(" 取消 ")]),_:1}),t(p,{secondary:"",round:"",type:"primary",loading:K.value,onClick:de},{default:s(()=>[c(" 激活 ")]),_:1},8,["loading"])])]),_:1})]),_:1})]),_:1},8,["model"])])):_("",!0)]),_:1})):_("",!0),t(F,{title:"账户安全",size:"small",class:"setting-card"},{default:s(()=>[c(" 您已设置密码 "),U.value?_("",!0):(r(),b(p,{key:0,quaternary:"",round:"",type:"success",onClick:e[14]||(e[14]=l=>U.value=!0)},{default:s(()=>[c(" 重置密码 ")]),_:1})),U.value?(r(),g("div",vt,[t(V,{ref_key:"formRef",ref:Q,model:a,rules:ve},{default:s(()=>[t(w,{path:"old_password",label:"旧密码"},{default:s(()=>[t(h,{value:a.old_password,"onUpdate:value":e[15]||(e[15]=l=>a.old_password=l),type:"password",placeholder:"请输入当前密码",onKeydown:e[16]||(e[16]=q(B(()=>{},["prevent"]),["enter"]))},null,8,["value"])]),_:1}),t(w,{path:"password",label:"新密码"},{default:s(()=>[t(h,{value:a.password,"onUpdate:value":e[17]||(e[17]=l=>a.password=l),type:"password",placeholder:"请输入新密码",onInput:re,onKeydown:e[18]||(e[18]=q(B(()=>{},["prevent"]),["enter"]))},null,8,["value"])]),_:1}),t(w,{ref_key:"rPasswordFormItemRef",ref:Y,first:"",path:"reenteredPassword",label:"重复密码"},{default:s(()=>[t(h,{value:a.reenteredPassword,"onUpdate:value":e[19]||(e[19]=l=>a.reenteredPassword=l),disabled:!a.password,type:"password",placeholder:"请再次输入密码",onKeydown:e[20]||(e[20]=q(B(()=>{},["prevent"]),["enter"]))},null,8,["value","disabled"])]),_:1},512),t(T,{gutter:[0,24]},{default:s(()=>[t(M,{span:24},{default:s(()=>[m("div",ft,[t(p,{quaternary:"",round:"",onClick:e[21]||(e[21]=l=>U.value=!1)},{default:s(()=>[c(" 取消 ")]),_:1}),t(p,{secondary:"",round:"",type:"primary",loading:D.value,onClick:ie},{default:s(()=>[c(" 更新 ")]),_:1},8,["loading"])])]),_:1})]),_:1})]),_:1},8,["model"])])):_("",!0)]),_:1})])}}});const jt=Ne(gt,[["__scopeId","data-v-a681720e"]]);export{jt as default};