import{W as Fe}from"./whisper-add-friend-B4hcy87Q.js";import{_ as Ne}from"./whisper-C90jedry.js";import{_ as Ge,a as Re}from"./post-item.vue_vue_type_style_index_0_lang-BMgNyw45.js";import{_ as qe}from"./post-skeleton-DOs8Pbo8.js";import{d as ye,r,c as le,b as ke,f as u,j as k,bq as h,k as D,q as a,w as l,t as A,Z as _,A as we,G as x,z as fe,F as ge,x as _e,U as Oe,C as Pe}from"./@vue-CT9fXtKS.js";import{u as Ae}from"./vuex-BtQJPySv.js";import{l as he}from"./lodash-CGvuAYz8.js";import{g as Se,a as Me,c as Ye,b as Le,d as We,e as Ke,u as je,f as Qe,h as He,_ as Ze}from"./index-CEQM3iaT.js";import{p as Je}from"./content-Doz2tKZo.js";import{V as R,P as X}from"./IEnum-B3rDUvtK.js";import{I as Xe,V as $e,A as et,d as tt,E as st}from"./@vicons-BcP-CCNH.js";import{n as be,t as at,i as nt,B as ze,v as ot,w as lt,x as it,y as rt,z as ut,A as ct,C as pt,a as Te,D as dt,E as vt,F as mt,G as ft,H as gt,k as _t,I as wt,J as ht,j as yt}from"./naive-ui-1pzJR9kY.js";import{_ as kt}from"./main-nav.vue_vue_type_style_index_0_lang-tx84Ctyj.js";import{b as At,u as bt}from"./vue-router-Oso6ZnuY.js";import{Y as zt}from"./v3-infinite-loading-DLisiGa3.js";import{S as Tt}from"./@opentiny-Brt0WuG3.js";import"./copy-to-clipboard-DgsYVcxl.js";import"./toggle-selection-DGa8lynz.js";import"./axios-upsvKRUO.js";import"./moment-P60zs0je.js";/* empty css               */import"./seemly-96Y8tuX4.js";import"./vueuc-rMdi6pp7.js";import"./evtd-CI_DDEu_.js";import"./@css-render-D-71Ub-V.js";import"./vooks-D483k878.js";import"./vdirs-Bxp-63WN.js";import"./@juggle-DY95s5UV.js";import"./css-render-wuhQizsj.js";import"./@emotion-WldOFDRm.js";import"./lodash-es-BI2Xm8S2.js";import"./treemate-N4GG0L_2.js";import"./async-validator-P8scd9xB.js";import"./date-fns-Db9XENWt.js";import"./paopao-video-player-DJIg2Ep1.js";import"./vue-BAiMcOs2.js";const Ct=$=>{const g=new FileReader,p=z=>["application/zip","application/x-zip","application/octet-stream","application/x-zip-compressed"].includes(z),w=()=>{const z=new Uint8Array(g.result).subarray(0,4);let d="";for(let b=0;b<z.length;b++)d+=z[b].toString(16);switch(d){case"504b0304":case"504b0506":case"504b0708":return p("application/zip");case"504b030414":return p("application/x-zip-compressed");case"504b0508":return p("application/x-zip");case"504b5370":return p("application/octet-stream");default:return!1}};return new Promise((z,d)=>{g.onloadend=()=>{const b=$.type;z(b===""||b==="application/octet-stream"?w():p(b))},g.readAsArrayBuffer($.slice(0,4))})},It={key:0,class:"compose-wrap"},Dt={class:"compose-line"},Bt={class:"compose-user"},Ut={class:"compose-line compose-options"},Vt={class:"attachment"},xt={class:"submit-wrap"},Et={class:"attachment-list-wrap"},Ft={key:0,class:"attachment-price-wrap"},Nt={key:0,class:"eye-wrap"},Gt={key:1,class:"link-wrap"},Rt={key:1,class:"compose-wrap"},qt={key:0,class:"login-only-wrap"},Ot={key:1,class:"login-wrap"},Pt=ye({__name:"compose",emits:["post-success"],setup($,{emit:g}){const p=g,w=Ae(),z=r([]),d=r(!1),b=r(!1),N=r(!1),G=r(!1),E=r(""),W=r([]),B=r(),q=r(0),y=r("public/image"),T=r([]),C=r([]),U=r([]),O=r([]),S=r(R.PUBLIC),v=r(R.PUBLIC),V=r("true".toLowerCase()==="true"),K="/v1/attachment",M=le(()=>"Bearer "+localStorage.getItem("PAOPAO_TOKEN")),Q=le(()=>{let t=[{value:R.PUBLIC,label:"公开"},{value:R.PRIVATE,label:"私密"},{value:R.Following,label:"关注可见"}];return w.state.profile.useFriendship&&t.push({value:R.FRIEND,label:"好友可见"}),t}),H=()=>{N.value=!N.value,N.value&&G.value&&(G.value=!1)},ee=()=>{G.value=!G.value,G.value&&N.value&&(N.value=!1)},te=he.debounce(t=>{Se({k:t}).then(e=>{let n=[];e.suggest.map(i=>{n.push({label:i,value:i})}),z.value=n,d.value=!1}).catch(e=>{d.value=!1})},200),ie=he.debounce(t=>{Me({k:t}).then(e=>{let n=[];e.suggest.map(i=>{n.push({label:i,value:i})}),z.value=n,d.value=!1}).catch(e=>{d.value=!1})},200),re=(t,e)=>{d.value||(d.value=!0,e==="@"?te(t):ie(t))},ue=t=>{t.length>w.state.profile.defaultTweetMaxLength?E.value=t.substring(0,w.state.profile.defaultTweetMaxLength):E.value=t},Z=t=>{y.value=t},se=t=>{for(let c=0;c<t.length;c++){var e=t[c].name,n=e.split(".").slice(0,-1).join("."),i=e.split(".").pop();n.length>30&&(t[c].name=n.substring(0,18)+"..."+n.substring(n.length-9)+"."+i)}T.value=t},ae=async t=>{var e,n,i,c,m;return y.value==="public/image"&&!["image/webp","image/png","image/jpg","image/jpeg","image/gif"].includes((e=t.file.file)==null?void 0:e.type)?(window.$message.warning("图片仅允许 webp/png/jpg/gif 格式"),!1):y.value==="image"&&((n=t.file.file)==null?void 0:n.size)>10485760?(window.$message.warning("图片大小不能超过10MB"),!1):y.value==="public/video"&&!["video/mp4","video/quicktime"].includes((i=t.file.file)==null?void 0:i.type)?(window.$message.warning("视频仅允许 mp4/mov 格式"),!1):y.value==="public/video"&&((c=t.file.file)==null?void 0:c.size)>104857600?(window.$message.warning("视频大小不能超过100MB"),!1):y.value==="attachment"&&!await Ct(t.file.file)?(window.$message.warning("附件仅允许 zip 格式"),!1):y.value==="attachment"&&((m=t.file.file)==null?void 0:m.size)>104857600?(window.$message.warning("附件大小不能超过100MB"),!1):!0},ne=({file:t,event:e})=>{var n;try{let i=JSON.parse((n=e.target)==null?void 0:n.response);i.code===0&&(y.value==="public/image"&&C.value.push({id:t.id,content:i.data.content}),y.value==="public/video"&&U.value.push({id:t.id,content:i.data.content}),y.value==="attachment"&&O.value.push({id:t.id,content:i.data.content}))}catch{window.$message.error("上传失败")}},oe=({file:t,event:e})=>{var n;try{let i=JSON.parse((n=e.target)==null?void 0:n.response);if(i.code!==0){let c=i.msg||"上传失败";i.details&&i.details.length>0&&i.details.map(m=>{c+=":"+m}),window.$message.error(c)}}catch{window.$message.error("上传失败")}},ce=({file:t})=>{let e=C.value.findIndex(n=>n.id===t.id);e>-1&&C.value.splice(e,1),e=U.value.findIndex(n=>n.id===t.id),e>-1&&U.value.splice(e,1),e=O.value.findIndex(n=>n.id===t.id),e>-1&&O.value.splice(e,1)},pe=()=>{if(E.value.trim().length===0){window.$message.warning("请输入内容哦");return}let{tags:t,users:e}=Je(E.value);const n=[];let i=100;n.push({content:E.value,type:X.TEXT,sort:i}),C.value.map(c=>{i++,n.push({content:c.content,type:X.IMAGEURL,sort:i})}),U.value.map(c=>{i++,n.push({content:c.content,type:X.VIDEOURL,sort:i})}),O.value.map(c=>{i++,n.push({content:c.content,type:X.ATTACHMENT,sort:i})}),W.value.length>0&&W.value.map(c=>{i++,n.push({content:c,type:X.LINKURL,sort:i})}),b.value=!0,Ye({contents:n,tags:Array.from(new Set(t)),users:Array.from(new Set(e)),attachment_price:+q.value*100,visibility:S.value}).then(c=>{var m;window.$message.success("发布成功"),b.value=!1,p("post-success",c),N.value=!1,G.value=!1,(m=B.value)==null||m.clear(),T.value=[],E.value="",W.value=[],C.value=[],U.value=[],O.value=[],S.value=v.value}).catch(c=>{b.value=!1})},Y=t=>{w.commit("triggerAuth",!0),w.commit("triggerAuthKey",t)};return ke(()=>{const t=w.state.profile.defaultTweetVisibility;w.state.profile.useFriendship&&t==="friend"?v.value=R.FRIEND:t==="following"?v.value=R.Following:t==="public"?v.value=R.PUBLIC:v.value=R.PRIVATE,S.value=v.value}),(t,e)=>{const n=be,i=at,c=nt,m=ze,J=ot,s=lt,o=it,P=rt,j=ut,F=ct,de=pt,ve=Te,me=dt,L=vt;return u(),k("div",null,[h(w).state.userInfo.id>0?(u(),k("div",It,[D("div",Dt,[D("div",Bt,[a(n,{round:"",size:30,src:h(w).state.userInfo.avatar},null,8,["src"])]),a(i,{type:"textarea",size:"large",autosize:"",bordered:!1,loading:d.value,value:E.value,prefix:["@","#"],options:z.value,onSearch:re,"onUpdate:value":ue,placeholder:"说说您的新鲜事..."},null,8,["loading","value","options"])]),a(F,{ref_key:"uploadRef",ref:B,abstract:"","list-type":"image",multiple:!0,max:9,action:K,headers:{Authorization:M.value},data:{type:y.value},"file-list":T.value,onBeforeUpload:ae,onFinish:ne,onError:oe,onRemove:ce,"onUpdate:fileList":se},{default:l(()=>[D("div",Ut,[D("div",Vt,[a(J,{abstract:""},{default:l(({handleClick:f})=>[a(m,{disabled:T.value.length>0&&y.value==="public/video"||T.value.length===9,onClick:()=>{Z("public/image"),f()},quaternary:"",circle:"",type:"primary"},{icon:l(()=>[a(c,{size:"20",color:"var(--primary-color)"},{default:l(()=>[a(h(Xe))]),_:1})]),_:2},1032,["disabled","onClick"])]),_:1}),h(w).state.profile.allowTweetVideo?(u(),A(J,{key:0,abstract:""},{default:l(({handleClick:f})=>[a(m,{disabled:T.value.length>0&&y.value!=="public/video"||T.value.length===9,onClick:()=>{Z("public/video"),f()},quaternary:"",circle:"",type:"primary"},{icon:l(()=>[a(c,{size:"20",color:"var(--primary-color)"},{default:l(()=>[a(h($e))]),_:1})]),_:2},1032,["disabled","onClick"])]),_:1})):_("",!0),h(w).state.profile.allowTweetAttachment?(u(),A(J,{key:1,abstract:""},{default:l(({handleClick:f})=>[a(m,{disabled:T.value.length>0&&y.value==="public/video"||T.value.length===9,onClick:()=>{Z("attachment"),f()},quaternary:"",circle:"",type:"primary"},{icon:l(()=>[a(c,{size:"20",color:"var(--primary-color)"},{default:l(()=>[a(h(et))]),_:1})]),_:2},1032,["disabled","onClick"])]),_:1})):_("",!0),a(m,{quaternary:"",circle:"",type:"primary",onClick:we(H,["stop"])},{icon:l(()=>[a(c,{size:"20",color:"var(--primary-color)"},{default:l(()=>[a(h(tt))]),_:1})]),_:1}),V.value?(u(),A(m,{key:2,quaternary:"",circle:"",type:"primary",onClick:we(ee,["stop"])},{icon:l(()=>[a(c,{size:"20",color:"var(--primary-color)"},{default:l(()=>[a(h(st))]),_:1})]),_:1})):_("",!0)]),D("div",xt,[a(o,{trigger:"hover",placement:"bottom"},{trigger:l(()=>[a(s,{class:"text-statistic",type:"circle","show-indicator":!1,status:"success","stroke-width":10,percentage:E.value.length/h(w).state.profile.defaultTweetMaxLength*100},null,8,["percentage"])]),default:l(()=>[x(" 已输入"+fe(E.value.length)+"字 ",1)]),_:1}),a(m,{loading:b.value,onClick:pe,type:"primary",secondary:"",round:""},{default:l(()=>e[6]||(e[6]=[x(" 发布 ")])),_:1},8,["loading"])])]),D("div",Et,[a(P),O.value.length>0?(u(),k("div",Ft,[h(w).state.profile.allowTweetAttachmentPrice?(u(),A(j,{key:0,value:q.value,"onUpdate:value":e[0]||(e[0]=f=>q.value=f),min:0,max:1e5,placeholder:"请输入附件价格，0为免费附件"},{prefix:l(()=>e[7]||(e[7]=[D("span",null," 附件价格￥",-1)])),_:1},8,["value"])):_("",!0)])):_("",!0)])]),_:1},8,["headers","data","file-list"]),G.value?(u(),k("div",Nt,[a(me,{value:S.value,"onUpdate:value":e[1]||(e[1]=f=>S.value=f),name:"radiogroup"},{default:l(()=>[a(ve,null,{default:l(()=>[(u(!0),k(ge,null,_e(Q.value,f=>(u(),A(de,{key:f.value,value:f.value,label:f.label},null,8,["value","label"]))),128))]),_:1})]),_:1},8,["value"])])):_("",!0),N.value?(u(),k("div",Gt,[a(L,{value:W.value,"onUpdate:value":e[2]||(e[2]=f=>W.value=f),placeholder:"请输入以http(s)://开头的链接",min:0,max:3},{"create-button-default":l(()=>e[8]||(e[8]=[x(" 创建链接 ")])),_:1},8,["value"])])):_("",!0)])):(u(),k("div",Rt,[e[12]||(e[12]=D("div",{class:"login-wrap"},[D("span",{class:"login-banner"}," 登录后，精彩更多")],-1)),h(w).state.profile.allowUserRegister?_("",!0):(u(),k("div",qt,[a(m,{strong:"",secondary:"",round:"",type:"primary",onClick:e[3]||(e[3]=f=>Y("signin"))},{default:l(()=>e[9]||(e[9]=[x(" 登录 ")])),_:1})])),h(w).state.profile.allowUserRegister?(u(),k("div",Ot,[a(m,{strong:"",secondary:"",round:"",type:"primary",onClick:e[4]||(e[4]=f=>Y("signin"))},{default:l(()=>e[10]||(e[10]=[x(" 登录 ")])),_:1}),a(m,{strong:"",secondary:"",round:"",type:"info",onClick:e[5]||(e[5]=f=>Y("signup"))},{default:l(()=>e[11]||(e[11]=[x(" 注册 ")])),_:1})])):_("",!0)]))])}}}),St="data:image/png;base64,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",Mt="/assets/discover-tweets-DGidPW73.jpeg",Yt="/assets/following-tweets-BKofJ8VU.jpeg",Lt={class:"slide-bar-item"},Wt={class:"slide-bar-item-title slide-bar-user-link"},Kt={key:1,class:"style-wrap"},jt={key:2,class:"skeleton-wrap"},Qt={key:0,class:"empty-wrap"},Ht={key:1},Zt={key:2},Jt={class:"load-more-wrap"},Xt={class:"load-more-spinner"},$t=ye({__name:"Home",setup($){const g=Ae(),p=At(),w=bt(),z=mt(),d=r("newest"),b=()=>{d.value="newest",t(B.value[0],0)},N=()=>{d.value="hots",t(B.value[1],1)},G=()=>{d.value="following",t(B.value[2],2)},E=r(9),W=r(8),B=r([{title:"最新动态",style:1,username:"",avatar:St,show:!0},{title:"热门推荐",style:2,username:"",avatar:Mt,show:!1},{title:"正在关注",style:3,username:"",avatar:Yt,show:!1},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0},{title:"",style:1,username:"",avatar:"",show:!0}]),q=Oe({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!1,is_following:!1,created_on:0,follows:0,followings:0,status:1}),y=r(null),T=r("泡泡广场"),C=r(!1),U=r(!1),O=r(1),S=r(""),v=r([]),V=r(1),K=r(20),M=r(0),Q=r(!1),H=r(!1),ee=r({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),te=s=>{ee.value=s,Q.value=!0},ie=()=>{Q.value=!1},re=()=>{H.value=!0},ue=s=>{z.warning({title:"删除好友",content:"将好友 “"+s.user.nickname+"” 删除，将同时删除 点赞/收藏 列表中关于该朋友的 “好友可见” 推文",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{He({user_id:q.id}).then(o=>{window.$message.success("操作成功"),s.user.is_friend=!1}).catch(o=>{})}})},Z=()=>{H.value=!1,y.value=null},se=s=>{y.value=s,q.id=s.user.id,q.username=s.user.username,q.nickname=s.user.nickname,s.user.is_friend?ue(s):re()},ae=s=>{z.success({title:"提示",content:"确定"+(s.user.is_following?"取消关注 @":"关注 @")+s.user.username+" 吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{s.user.is_following?je({user_id:s.user.id}).then(o=>{window.$message.success("操作成功"),ne(s.user_id,!1)}).catch(o=>{}):Qe({user_id:s.user.id}).then(o=>{window.$message.success("关注成功"),ne(s.user_id,!0)}).catch(o=>{})}})};function ne(s,o){for(let P in v.value)v.value[P].user_id==s&&(v.value[P].user.is_following=o)}const oe=()=>{T.value="泡泡广场",p.query&&p.query.q&&(p.query.t&&p.query.t==="tag"?T.value="#"+decodeURIComponent(p.query.q):T.value="搜索: "+decodeURIComponent(p.query.q))},ce=le(()=>g.state.userInfo.id>0&&!g.state.profile.enableTrendsBar&&g.state.desktopModelShow),pe=le(()=>g.state.profile.useFriendship&&g.state.profile.enableTrendsBar&&g.state.desktopModelShow&&g.state.userInfo.id>0),Y=()=>{C.value=!1,U.value=!1,v.value=[],V.value=1,M.value=0},t=(s,o)=>{switch(Y(),O.value=s.style,p.query.q&&(p.query.q=null,oe()),s.style){case 1:n("newest");break;case 2:n("hots");break;case 3:p.query.q=null,n("following");break;case 21:S.value=s.username,i();break}B.value[o].show=!1},e=()=>{B.value=B.value.slice(0,3),!(!g.state.profile.useFriendship||!g.state.profile.enableTrendsBar||g.state.userInfo.id===0)&&Le({page:1,page_size:50}).then(s=>{var o=0;const P=s.list||[];let j=[];for(;o<P.length;o++){let F=P[o];j.push({title:F.nickname,style:21,username:F.username,avatar:F.avatar,show:F.is_fresh})}j.length>0&&(B.value=B.value.concat(j))}).catch(s=>{console.log(s)})},n=s=>{C.value=!0,We({query:p.query.q?decodeURIComponent(p.query.q):null,type:p.query.t,style:s,page:V.value,page_size:K.value}).then(o=>{C.value=!1,o.list.length===0&&(U.value=!0),V.value>1?v.value=v.value.concat(o.list):(v.value=o.list,window.scrollTo(0,0)),M.value=Math.ceil(o.pager.total_rows/K.value)}).catch(o=>{C.value=!1,V.value>1&&V.value--})},i=()=>{C.value=!0,Ke({username:S.value,style:"post",page:V.value,page_size:K.value}).then(s=>{C.value=!1,s.list.length===0&&(U.value=!0),V.value>1?v.value=v.value.concat(s.list):(v.value=s.list||[],window.scrollTo(0,0)),M.value=Math.ceil(s.pager.total_rows/K.value)}).catch(s=>{v.value=[],V.value>1&&V.value--,C.value=!1})},c=s=>{w.push({name:"post",query:{id:s.id}})},m=()=>{switch(O.value){case 1:n("newest");break;case 2:n("hots");break;case 3:n("following");break;case 21:p.query.q?n("search"):i();break}},J=()=>{V.value<M.value||M.value==0?(U.value=!1,V.value++,m()):U.value=!0};return ke(()=>{Y(),e(),n("newest")}),Pe(()=>({path:p.path,query:p.query,refresh:g.state.refresh}),(s,o)=>{if(oe(),s.refresh!==o.refresh){Y(),setTimeout(()=>{e(),m()},0);return}o.path!=="/post"&&s.path==="/"&&(Y(),setTimeout(()=>{e(),m()},0))}),(s,o)=>{const P=kt,j=Pt,F=gt,de=be,ve=_t,me=wt,L=ze,f=Te,Ce=qe,Ie=ht,De=Ge,Be=Re,Ue=Ne,Ve=Fe,xe=ft,Ee=yt;return u(),k("div",null,[a(P,{title:T.value},null,8,["title"]),a(xe,{class:"main-content-wrap",bordered:""},{default:l(()=>[a(F,null,{default:l(()=>[a(j,{onPostSuccess:c})]),_:1}),pe.value?(u(),A(F,{key:0},{default:l(()=>[a(h(Tt),{modelValue:B.value,"onUpdate:modelValue":o[0]||(o[0]=I=>B.value=I),"wheel-blocks":W.value,"init-blocks":E.value,onClick:t,tag:"div","sub-tag":"div"},{default:l(I=>[D("div",Lt,[a(ve,{value:"1",offset:[-4,48],dot:"",show:I.slotData.show},{default:l(()=>[a(de,{round:"",size:48,src:I.slotData.avatar,class:"slide-bar-item-avatar"},null,8,["src"])]),_:2},1032,["show"]),D("div",Wt,[a(me,{"line-clamp":2},{default:l(()=>[x(fe(I.slotData.title),1)]),_:2},1024)])])]),_:1},8,["modelValue","wheel-blocks","init-blocks"])]),_:1})):ce.value?(u(),k("div",Kt,[a(f,null,{default:l(()=>[d.value!=="newest"?(u(),A(L,{key:0,size:"small",bordered:!1,onClick:b,class:"style-item",secondary:"",round:""},{default:l(()=>o[2]||(o[2]=[x(" 全部 ")])),_:1})):_("",!0),d.value==="newest"?(u(),A(L,{key:1,size:"small",type:"success",bordered:!1,onClick:b,class:"style-item",secondary:"",round:""},{default:l(()=>o[3]||(o[3]=[x(" 全部 ")])),_:1})):_("",!0),d.value!=="hots"?(u(),A(L,{key:2,size:"small",bordered:!1,onClick:N,class:"style-item",secondary:"",round:""},{default:l(()=>o[4]||(o[4]=[x(" 热门推荐 ")])),_:1})):_("",!0),d.value==="hots"?(u(),A(L,{key:3,size:"small",type:"success",bordered:!1,onClick:N,class:"style-item",secondary:"",round:""},{default:l(()=>o[5]||(o[5]=[x(" 热门推荐 ")])),_:1})):_("",!0),d.value!=="following"?(u(),A(L,{key:4,size:"small",bordered:!1,onClick:G,class:"style-item",secondary:"",round:""},{default:l(()=>o[6]||(o[6]=[x(" 正在关注 ")])),_:1})):_("",!0),d.value==="following"?(u(),A(L,{key:5,size:"small",type:"success",bordered:!1,onClick:G,class:"style-item",secondary:"",round:""},{default:l(()=>o[7]||(o[7]=[x(" 正在关注 ")])),_:1})):_("",!0)]),_:1})])):_("",!0),C.value&&v.value.length===0?(u(),k("div",jt,[a(Ce,{num:K.value},null,8,["num"])])):_("",!0),D("div",null,[v.value.length===0?(u(),k("div",Qt,[a(Ie,{size:"large",description:"暂无数据"})])):_("",!0),h(g).state.desktopModelShow?(u(),k("div",Ht,[(u(!0),k(ge,null,_e(v.value,I=>(u(),A(F,{key:I.id},{default:l(()=>[a(De,{post:I,isOwner:h(g).state.userInfo.id==I.user_id,addFollowAction:!0,onSendWhisper:te,onHandleFollowAction:ae,onHandleFriendAction:se},null,8,["post","isOwner"])]),_:2},1024))),128))])):(u(),k("div",Zt,[(u(!0),k(ge,null,_e(v.value,I=>(u(),A(F,{key:I.id},{default:l(()=>[a(Be,{post:I,isOwner:h(g).state.userInfo.id==I.user_id,addFollowAction:!0,onSendWhisper:te,onHandleFollowAction:ae,onHandleFriendAction:se},null,8,["post","isOwner"])]),_:2},1024))),128))]))]),a(Ue,{show:Q.value,user:ee.value,onSuccess:ie},null,8,["show","user"]),a(Ve,{show:H.value,user:q,onSuccess:Z},null,8,["show","user"])]),_:1}),M.value>0?(u(),A(f,{key:0,justify:"center"},{default:l(()=>[a(h(zt),{class:"load-more",slots:{complete:"没有更多泡泡了",error:"加载出错"},onInfinite:o[1]||(o[1]=I=>J())},{spinner:l(()=>[D("div",Jt,[U.value?_("",!0):(u(),A(Ee,{key:0,size:14})),D("span",Xt,fe(U.value?"没有更多泡泡了":"加载更多"),1)])]),_:1})]),_:1})):_("",!0)])}}}),Gs=Ze($t,[["__scopeId","data-v-cc7d12d2"]]);export{Gs as default};
