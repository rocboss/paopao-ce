import{_ as Fe,a as Me}from"./post-item.vue_vue_type_style_index_0_lang-25504051.js";import{_ as qe}from"./post-skeleton-d2e839d6.js";import{C as Se,_ as V,D as We,E as Oe,B as S,F as Be,G as Re,H as De,I as Ae}from"./index-152c5794.js";import{R as se,K as ae,S as te,b as oe,e as H,i as ne,T as Ne,F as Ve,G as He,j as Q,o as Ee,O as Ge,P as je,k as Le,f as Ke,g as Qe,I as Je,H as Xe}from"./naive-ui-62663ad7.js";import{d as E,r as i,o as u,Q as U,a1 as o,a as r,V as a,e as g,M as w,E as Ye,n as Ze,w as es,j as ss,a3 as as,c as f,_ as h,O as P,a2 as J,F as X,a4 as Y,s as Z}from"./@vue-e0e89260.js";import{_ as ts}from"./main-nav.vue_vue_type_style_index_0_lang-32d416c3.js";import{u as os}from"./vuex-473b3783.js";import{b as ns}from"./vue-router-b8e3382f.js";import{b as ls}from"./formatTime-4210fcd1.js";import{i as is,w as us,x as ee,y as cs,z as rs,D as _s,G as ps}from"./@vicons-0524c43e.js";import"./content-ff44e340.js";import"./paopao-video-player-aa5e8b3f.js";import"./copy-to-clipboard-1dd3075d.js";import"./toggle-selection-93f4ad84.js";import"./axios-4a70c6fc.js";/* empty css               */import"./seemly-76b7b838.js";import"./vueuc-59ca65c3.js";import"./evtd-b614532e.js";import"./@css-render-580d83ec.js";import"./vooks-a50491fd.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";import"./moment-2ab8298d.js";const ds={class:"whisper-wrap"},ms={class:"whisper-line"},vs={class:"whisper-line send-wrap"},fs=E({__name:"whisper",props:{show:{type:Boolean,default:!1},user:{}},emits:["success"],setup(W,{emit:b}){const d=W,_=i(""),t=i(!1),s=()=>{b("success")},l=()=>{t.value=!0,Se({user_id:d.user.id,content:_.value}).then(p=>{window.$message.success("发送成功"),t.value=!1,_.value="",s()}).catch(p=>{t.value=!1})};return(p,k)=>{const c=se,m=ae,n=te,y=oe,z=H,I=ne;return u(),U(I,{show:p.show,"onUpdate:show":s,class:"whisper-card",preset:"card",size:"small",title:"私信","mask-closable":!1,bordered:!1,style:{width:"360px"}},{default:o(()=>[r("div",ds,[a(n,{"show-icon":!1},{default:o(()=>[g(" 即将发送私信给: "),a(m,{style:{"max-width":"100%"}},{default:o(()=>[a(c,{type:"success"},{default:o(()=>[g(w(p.user.nickname)+"@"+w(p.user.username),1)]),_:1})]),_:1})]),_:1}),r("div",ms,[a(y,{type:"textarea",placeholder:"请输入私信内容（请勿发送不和谐内容，否则将会被封号）",autosize:{minRows:5,maxRows:10},value:_.value,"onUpdate:value":k[0]||(k[0]=T=>_.value=T),maxlength:"200","show-count":""},null,8,["value"])]),r("div",vs,[a(z,{strong:"",secondary:"",type:"primary",loading:t.value,onClick:l},{default:o(()=>[g(" 发送 ")]),_:1},8,["loading"])])])]),_:1},8,["show"])}}});const hs=V(fs,[["__scopeId","data-v-0cbfe47c"]]),gs={class:"whisper-wrap"},ws={class:"whisper-line"},ks={class:"whisper-line send-wrap"},ys=E({__name:"whisper-add-friend",props:{show:{type:Boolean,default:!1},user:{}},emits:["success"],setup(W,{emit:b}){const d=W,_=i(""),t=i(!1),s=()=>{b("success")},l=()=>{t.value=!0,We({user_id:d.user.id,greetings:_.value}).then(p=>{window.$message.success("发送成功"),t.value=!1,_.value="",s()}).catch(p=>{t.value=!1})};return(p,k)=>{const c=se,m=ae,n=te,y=oe,z=H,I=ne;return u(),U(I,{show:p.show,"onUpdate:show":s,class:"whisper-card",preset:"card",size:"small",title:"申请添加朋友","mask-closable":!1,bordered:!1,style:{width:"360px"}},{default:o(()=>[r("div",gs,[a(n,{"show-icon":!1},{default:o(()=>[g(" 发送添加朋友申请给: "),a(m,{style:{"max-width":"100%"}},{default:o(()=>[a(c,{type:"success"},{default:o(()=>[g(w(p.user.nickname)+"@"+w(p.user.username),1)]),_:1})]),_:1})]),_:1}),r("div",ws,[a(y,{type:"textarea",placeholder:"请输入真挚的问候语",autosize:{minRows:5,maxRows:10},value:_.value,"onUpdate:value":k[0]||(k[0]=T=>_.value=T),maxlength:"120","show-count":""},null,8,["value"])]),r("div",ks,[a(z,{strong:"",secondary:"",type:"primary",loading:t.value,onClick:l},{default:o(()=>[g(" 发送 ")]),_:1},8,["loading"])])])]),_:1},8,["show"])}}});const bs=V(ys,[["__scopeId","data-v-60be56a2"]]),$s={key:0,class:"profile-baseinfo"},xs={class:"avatar"},Ps={class:"base-info"},Us={class:"username"},zs={class:"userinfo"},Is={class:"info-item"},Ts={class:"info-item"},Cs={class:"userinfo"},Fs={class:"info-item"},Ms={class:"info-item"},qs={key:0,class:"user-opts"},Ss={key:0,class:"skeleton-wrap"},Ws={key:1},Os={key:0,class:"empty-wrap"},Bs={key:1},Rs={key:2},Ds={key:0,class:"pagination-wrap"},As=E({__name:"User",setup(W){const b=Ne(),d=os(),_=ns(),t=i(!1),s=Ye({id:0,avatar:"",username:"",nickname:"",is_admin:!1,is_friend:!0,is_following:!1,created_on:0,follows:0,followings:0,status:1}),l=i(!1),p=i(!1),k=i(!1),c=i([]),m=i(_.query.s||""),n=i(+_.query.p||1),y=i("post"),z=i(+_.query.p||1),I=i(1),T=i(1),G=i(1),j=i(1),v=i(20),$=i(0),le=()=>{switch(y.value){case"post":O();break;case"comment":B();break;case"highlight":R();break;case"media":D();break;case"star":A();break}},O=()=>{t.value=!0,S({username:m.value,style:"post",page:n.value,page_size:v.value}).then(e=>{t.value=!1,c.value=e.list||[],$.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{c.value=[],t.value=!1})},B=()=>{t.value=!0,S({username:m.value,style:"comment",page:n.value,page_size:v.value}).then(e=>{t.value=!1,c.value=e.list||[],$.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{c.value=[],t.value=!1})},R=()=>{t.value=!0,S({username:m.value,style:"highlight",page:n.value,page_size:v.value}).then(e=>{t.value=!1,c.value=e.list||[],$.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{c.value=[],t.value=!1})},D=()=>{t.value=!0,S({username:m.value,style:"media",page:n.value,page_size:v.value}).then(e=>{t.value=!1,c.value=e.list||[],$.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{c.value=[],t.value=!1})},A=()=>{t.value=!0,S({username:m.value,style:"star",page:n.value,page_size:v.value}).then(e=>{t.value=!1,c.value=e.list||[],$.value=Math.ceil(e.pager.total_rows/v.value),window.scrollTo(0,0)}).catch(e=>{c.value=[],t.value=!1})},ie=e=>{switch(y.value=e,y.value){case"post":n.value=z.value,O();break;case"comment":n.value=I.value,B();break;case"highlight":n.value=T.value,R();break;case"media":n.value=G.value,D();break;case"star":n.value=j.value,A();break}},F=()=>{l.value=!0,Oe({username:m.value}).then(e=>{l.value=!1,s.id=e.id,s.avatar=e.avatar,s.username=e.username,s.nickname=e.nickname,s.is_admin=e.is_admin,s.is_friend=e.is_friend,s.created_on=e.created_on,s.is_following=e.is_following,s.follows=e.follows,s.followings=e.followings,s.status=e.status,le()}).catch(e=>{l.value=!1,console.log(e)})},ue=e=>{switch(n.value=e,y.value){case"post":z.value=e,O();break;case"comment":I.value=n.value,B();break;case"highlight":T.value=n.value,R();break;case"media":G.value=n.value,D();break;case"star":j.value=n.value,A();break}},ce=()=>{p.value=!0},re=()=>{k.value=!0},_e=()=>{p.value=!1},pe=()=>{k.value=!1},x=e=>()=>Z(Q,null,{default:()=>Z(e)}),de=Ze(()=>{let e=[{label:"私信",key:"whisper",icon:x(us)}];return d.state.userInfo.is_admin&&(s.status===1?e.push({label:"禁言",key:"banned",icon:x(ee)}):e.push({label:"解封",key:"deblocking",icon:x(ee)})),s.is_following?e.push({label:"取消关注",key:"unfollow",icon:x(cs)}):e.push({label:"关注",key:"follow",icon:x(rs)}),s.is_friend?e.push({label:"删除好友",key:"delete",icon:x(_s)}):e.push({label:"添加朋友",key:"requesting",icon:x(ps)}),e}),me=e=>{switch(e){case"whisper":ce();break;case"delete":ve();break;case"requesting":re();break;case"follow":case"unfollow":fe();break;case"banned":case"deblocking":he();break}},ve=()=>{b.warning({title:"删除好友",content:"将好友 “"+s.nickname+"” 删除，将同时删除 点赞/收藏 列表中关于该朋友的 “好友可见” 推文",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{l.value=!0,Be({user_id:s.id}).then(e=>{l.value=!1,s.is_friend=!1,O()}).catch(e=>{l.value=!1,console.log(e)})}})},fe=()=>{b.success({title:"提示",content:"确定"+(s.is_following?"取消关注":"关注")+"该用户吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{l.value=!0,s.is_following?Re({user_id:s.id}).then(e=>{l.value=!1,window.$message.success("取消关注成功"),F()}).catch(e=>{l.value=!1,console.log(e)}):De({user_id:s.id}).then(e=>{l.value=!1,window.$message.success("关注成功"),F()}).catch(e=>{l.value=!1,console.log(e)})}})},he=()=>{b.warning({title:"警告",content:"确定对该用户进行"+(s.status===1?"禁言":"解封")+"处理吗？",positiveText:"确定",negativeText:"取消",onPositiveClick:()=>{l.value=!0,Ae({id:s.id,status:s.status===1?2:1}).then(e=>{l.value=!1,s.status===1?window.$message.success("禁言成功"):window.$message.success("解封成功"),F()}).catch(e=>{l.value=!1,console.log(e)})}})};return es(()=>({path:_.path,query:_.query}),(e,C)=>{C.path==="/user"&&e.path==="/user"&&(m.value=_.query.s||"",F())}),ss(()=>{F()}),(e,C)=>{const ge=ts,we=Ee,N=Ge,L=as("router-link"),ke=H,ye=je,be=hs,$e=Le,M=Ke,xe=Qe,Pe=qe,Ue=Je,ze=Fe,K=Xe,Ie=Me,Te=Ve,Ce=He;return u(),f("div",null,[a(ge,{title:"用户详情"}),a(Te,{class:"main-content-wrap profile-wrap",bordered:""},{default:o(()=>[a($e,{show:l.value},{default:o(()=>[s.id>0?(u(),f("div",$s,[r("div",xs,[a(we,{size:72,src:s.avatar},null,8,["src"])]),r("div",Ps,[r("div",Us,[r("strong",null,w(s.nickname),1),r("span",null," @"+w(s.username),1),h(d).state.userInfo.id>0&&h(d).state.userInfo.username!=s.username&&s.is_friend?(u(),U(N,{key:0,class:"top-tag",type:"info",size:"small",round:""},{default:o(()=>[g(" 好友 ")]),_:1})):P("",!0),h(d).state.userInfo.id>0&&h(d).state.userInfo.username!=s.username&&s.is_following?(u(),U(N,{key:1,class:"top-tag",type:"success",size:"small",round:""},{default:o(()=>[g(" 已关注 ")]),_:1})):P("",!0),s.is_admin?(u(),U(N,{key:2,class:"top-tag",type:"error",size:"small",round:""},{default:o(()=>[g(" 管理员 ")]),_:1})):P("",!0)]),r("div",zs,[r("span",Is,"UID. "+w(s.id),1),r("span",Ts,w(h(ls)(s.created_on))+" 加入",1)]),r("div",Cs,[r("span",Fs,[a(L,{onClick:C[0]||(C[0]=J(()=>{},["stop"])),class:"following-link",to:{name:"following",query:{s:s.username,n:s.nickname,t:"follows"}}},{default:o(()=>[g(" 关注  "+w(s.follows),1)]),_:1},8,["to"])]),r("span",Ms,[a(L,{onClick:C[1]||(C[1]=J(()=>{},["stop"])),class:"following-link",to:{name:"following",query:{s:s.username,n:s.nickname,t:"followings"}}},{default:o(()=>[g(" 粉丝  "+w(s.followings),1)]),_:1},8,["to"])])])]),h(d).state.userInfo.id>0&&h(d).state.userInfo.username!=s.username?(u(),f("div",qs,[a(ye,{placement:"bottom-end",trigger:"click",size:"small",options:de.value,onSelect:me},{default:o(()=>[a(ke,{quaternary:"",circle:""},{icon:o(()=>[a(h(Q),null,{default:o(()=>[a(h(is))]),_:1})]),_:1})]),_:1},8,["options"])])):P("",!0)])):P("",!0),a(be,{show:p.value,user:s,onSuccess:_e},null,8,["show","user"]),a(bs,{show:k.value,user:s,onSuccess:pe},null,8,["show","user"])]),_:1},8,["show"]),a(xe,{class:"profile-tabs-wrap",type:"line",animated:"","onUpdate:value":ie},{default:o(()=>[a(M,{name:"post",tab:"泡泡"}),a(M,{name:"comment",tab:"评论"}),a(M,{name:"highlight",tab:"亮点"}),a(M,{name:"media",tab:"图文"}),a(M,{name:"star",tab:"喜欢"})]),_:1}),t.value?(u(),f("div",Ss,[a(Pe,{num:v.value},null,8,["num"])])):(u(),f("div",Ws,[c.value.length===0?(u(),f("div",Os,[a(Ue,{size:"large",description:"暂无数据"})])):P("",!0),h(d).state.desktopModelShow?(u(),f("div",Bs,[(u(!0),f(X,null,Y(c.value,q=>(u(),U(K,{key:q.id},{default:o(()=>[a(ze,{post:q},null,8,["post"])]),_:2},1024))),128))])):(u(),f("div",Rs,[(u(!0),f(X,null,Y(c.value,q=>(u(),U(K,{key:q.id},{default:o(()=>[a(Ie,{post:q},null,8,["post"])]),_:2},1024))),128))]))]))]),_:1}),$.value>0?(u(),f("div",Ds,[a(Ce,{page:n.value,"onUpdate:page":ue,"page-slot":h(d).state.collapsedRight?5:8,"page-count":$.value},null,8,["page","page-slot","page-count"])])):P("",!0)])}}});const ha=V(As,[["__scopeId","data-v-42bb9831"]]);export{ha as default};
