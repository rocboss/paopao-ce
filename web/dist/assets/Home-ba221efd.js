import{_ as xe,a as ze}from"./post-item.vue_vue_type_style_index_0_lang-ebd1ae42.js";import{_ as Ie}from"./post-skeleton-3703f541.js";import{d as te,r as l,j as ae,o as u,c as h,_ as z,a as T,V as a,a2 as s,Q as R,O as C,a3 as Y,e as L,M as Q,F as Z,a5 as W,n as qe,w as Ue}from"./@vue-7e1ab0af.js";import{u as oe}from"./vuex-f1ee712f.js";import{l as ee}from"./lodash-e0b37ac3.js";import{g as Ae,a as Ee,c as Re,b as Pe,_ as Se}from"./index-b4b0f710.js";import{p as Le}from"./content-1c30deb5.js";import{V as E,P as V}from"./IEnum-a180d93e.js";import{I as Ve,V as Ne,A as Fe,d as Be,E as Me}from"./@vicons-b98681e0.js";import{o as Oe,v as je,j as De,e as He,w as Ge,x as Ke,y as Je,z as Qe,A as Ze,B as We,C as Xe,a as ne,D as Ye,E as et,F as tt,G as at,H as ot,k as nt}from"./naive-ui-702193c2.js";import{_ as st}from"./main-nav.vue_vue_type_style_index_0_lang-fa3b58e7.js";import{b as lt,u as it}from"./vue-router-edf90322.js";import{W as rt}from"./v3-infinite-loading-0611cb6c.js";import"./formatTime-4210fcd1.js";import"./moment-2ab8298d.js";import"./copy-to-clipboard-4ef7d3eb.js";import"./@babel-725317a4.js";import"./toggle-selection-93f4ad84.js";import"./axios-4a70c6fc.js";/* empty css */import"./seemly-76b7b838.js";import"./vueuc-2fc92f18.js";import"./evtd-b614532e.js";import"./@css-render-16be7445.js";import"./vooks-e23078ea.js";import"./vdirs-b0483831.js";import"./@juggle-41516555.js";import"./css-render-6a5c5852.js";import"./@emotion-8a8e73f6.js";import"./lodash-es-8412e618.js";import"./treemate-25c27bff.js";import"./async-validator-dee29e8b.js";import"./date-fns-975a2d8f.js";import"./paopao-video-player-66a1a537.js";const ut=N=>{const q=new FileReader,c=_=>["application/zip","application/x-zip","application/octet-stream","application/x-zip-compressed"].includes(_),P=()=>{const _=new Uint8Array(q.result).subarray(0,4);let m="";for(let i=0;i<_.length;i++)m+=_[i].toString(16);switch(m){case"504b0304":case"504b0506":case"504b0708":return c("application/zip");case"504b030414":return c("application/x-zip-compressed");case"504b0508":return c("application/x-zip");case"504b5370":return c("application/octet-stream");default:return!1}};return new Promise((_,m)=>{q.onloadend=()=>{const i=N.type;_(i===""||i==="application/octet-stream"?P():c(i))},q.readAsArrayBuffer(N.slice(0,4))})},ct={key:0,class:"compose-wrap"},pt={class:"compose-line"},_t={class:"compose-user"},dt={class:"compose-line compose-options"},mt={class:"attachment"},vt={class:"submit-wrap"},ft={class:"attachment-list-wrap"},gt={key:0,class:"attachment-price-wrap"},yt=T("span",null," 附件价格￥",-1),ht={key:0,class:"eye-wrap"},wt={key:1,class:"link-wrap"},bt={key:1,class:"compose-wrap"},kt=T("div",{class:"login-wrap"},[T("span",{class:"login-banner"}," 登录后，精彩更多")],-1),Ct={key:0,class:"login-only-wrap"},$t={key:1,class:"login-wrap"},Tt=te({__name:"compose",emits:["post-success"],setup(N,{emit:q}){const c=oe(),P=l([]),_=l(!1),m=l(!1),i=l(!1),v=l(!1),w=l(""),I=l([]),F=l(),U=l(0),g=l("public/image"),x=l([]),p=l([]),y=l([]),k=l([]),b=l(E.FRIEND),$=l(E.FRIEND),M=[{value:E.PUBLIC,label:"公开"},{value:E.PRIVATE,label:"私密"},{value:E.FRIEND,label:"好友可见"}],S=+"400",B=l("true".toLowerCase()==="true"),O=l("true".toLowerCase()==="true"),j=l("true".toLowerCase()==="true"),D=l("false".toLowerCase()==="true"),H=l("true".toLowerCase()==="true"),A="/v1/attachment",X=l(),se=()=>{i.value=!i.value,i.value&&v.value&&(v.value=!1)},le=()=>{v.value=!v.value,v.value&&i.value&&(i.value=!1)},ie=ee.debounce(t=>{Ae({k:t}).then(e=>{let n=[];e.suggest.map(o=>{n.push({label:o,value:o})}),P.value=n,_.value=!1}).catch(e=>{_.value=!1})},200),re=ee.debounce(t=>{Ee({k:t}).then(e=>{let n=[];e.suggest.map(o=>{n.push({label:o,value:o})}),P.value=n,_.value=!1}).catch(e=>{_.value=!1})},200),ue=(t,e)=>{_.value||(_.value=!0,e==="@"?ie(t):re(t))},ce=t=>{t.length>S?w.value=t.substring(0,S):w.value=t},G=t=>{g.value=t},pe=t=>{for(let r=0;r<t.length;r++){var e=t[r].name,n=e.split(".").slice(0,-1).join("."),o=e.split(".").pop();n.length>30&&(t[r].name=n.substring(0,18)+"..."+n.substring(n.length-9)+"."+o)}x.value=t},_e=async t=>{var e,n,o,r,f;return g.value==="public/image"&&!["image/png","image/jpg","image/jpeg","image/gif"].includes((e=t.file.file)==null?void 0:e.type)?(window.$message.warning("图片仅允许 png/jpg/gif 格式"),!1):g.value==="image"&&((n=t.file.file)==null?void 0:n.size)>10485760?(window.$message.warning("图片大小不能超过10MB"),!1):g.value==="public/video"&&!["video/mp4","video/quicktime"].includes((o=t.file.file)==null?void 0:o.type)?(window.$message.warning("视频仅允许 mp4/mov 格式"),!1):g.value==="public/video"&&((r=t.file.file)==null?void 0:r.size)>104857600?(window.$message.warning("视频大小不能超过100MB"),!1):g.value==="attachment"&&!await ut(t.file.file)?(window.$message.warning("附件仅允许 zip 格式"),!1):g.value==="attachment"&&((f=t.file.file)==null?void 0:f.size)>104857600?(window.$message.warning("附件大小不能超过100MB"),!1):!0},de=({file:t,event:e})=>{var n;try{let o=JSON.parse((n=e.target)==null?void 0:n.response);o.code===0&&(g.value==="public/image"&&p.value.push({id:t.id,content:o.data.content}),g.value==="public/video"&&y.value.push({id:t.id,content:o.data.content}),g.value==="attachment"&&k.value.push({id:t.id,content:o.data.content}))}catch{window.$message.error("上传失败")}},me=({file:t,event:e})=>{var n;try{let o=JSON.parse((n=e.target)==null?void 0:n.response);if(o.code!==0){let r=o.msg||"上传失败";o.details&&o.details.length>0&&o.details.map(f=>{r+=":"+f}),window.$message.error(r)}}catch{window.$message.error("上传失败")}},ve=({file:t})=>{let e=p.value.findIndex(n=>n.id===t.id);e>-1&&p.value.splice(e,1),e=y.value.findIndex(n=>n.id===t.id),e>-1&&y.value.splice(e,1),e=k.value.findIndex(n=>n.id===t.id),e>-1&&k.value.splice(e,1)},fe=()=>{if(w.value.trim().length===0){window.$message.warning("请输入内容哦");return}let{tags:t,users:e}=Le(w.value);const n=[];let o=100;n.push({content:w.value,type:V.TEXT,sort:o}),p.value.map(r=>{o++,n.push({content:r.content,type:V.IMAGEURL,sort:o})}),y.value.map(r=>{o++,n.push({content:r.content,type:V.VIDEOURL,sort:o})}),k.value.map(r=>{o++,n.push({content:r.content,type:V.ATTACHMENT,sort:o})}),I.value.length>0&&I.value.map(r=>{o++,n.push({content:r,type:V.LINKURL,sort:o})}),m.value=!0,Re({contents:n,tags:Array.from(new Set(t)),users:Array.from(new Set(e)),attachment_price:+U.value*100,visibility:b.value}).then(r=>{var f;window.$message.success("发布成功"),m.value=!1,q("post-success",r),i.value=!1,v.value=!1,(f=F.value)==null||f.clear(),x.value=[],w.value="",I.value=[],p.value=[],y.value=[],k.value=[],b.value=$.value}).catch(r=>{m.value=!1})},K=t=>{c.commit("triggerAuth",!0),c.commit("triggerAuthKey",t)};return ae(()=>{"friend".toLowerCase()==="friend"?$.value=E.FRIEND:"friend".toLowerCase()==="public"?$.value=E.PUBLIC:$.value=E.PRIVATE,b.value=$.value,X.value="Bearer "+localStorage.getItem("PAOPAO_TOKEN")}),(t,e)=>{const n=Oe,o=je,r=De,f=He,J=Ge,ge=Ke,ye=Je,he=Qe,we=Ze,be=We,ke=Xe,Ce=ne,$e=Ye,Te=et;return u(),h("div",null,[z(c).state.userInfo.id>0?(u(),h("div",ct,[T("div",pt,[T("div",_t,[a(n,{round:"",size:30,src:z(c).state.userInfo.avatar},null,8,["src"])]),a(o,{type:"textarea",size:"large",autosize:"",bordered:!1,loading:_.value,value:w.value,prefix:["@","#"],options:P.value,onSearch:ue,"onUpdate:value":ce,placeholder:"说说您的新鲜事..."},null,8,["loading","value","options"])]),a(be,{ref_key:"uploadRef",ref:F,abstract:"","list-type":"image",multiple:!0,max:9,action:A,headers:{Authorization:X.value},data:{type:g.value},"file-list":x.value,onBeforeUpload:_e,onFinish:de,onError:me,onRemove:ve,"onUpdate:fileList":pe},{default:s(()=>[T("div",dt,[T("div",mt,[a(J,{abstract:""},{default:s(({handleClick:d})=>[a(f,{disabled:x.value.length>0&&g.value==="public/video"||x.value.length===9,onClick:()=>{G("public/image"),d()},quaternary:"",circle:"",type:"primary"},{icon:s(()=>[a(r,{size:"20",color:"var(--primary-color)"},{default:s(()=>[a(z(Ve))]),_:1})]),_:2},1032,["disabled","onClick"])]),_:1}),O.value?(u(),R(J,{key:0,abstract:""},{default:s(({handleClick:d})=>[a(f,{disabled:x.value.length>0&&g.value!=="public/video"||x.value.length===9,onClick:()=>{G("public/video"),d()},quaternary:"",circle:"",type:"primary"},{icon:s(()=>[a(r,{size:"20",color:"var(--primary-color)"},{default:s(()=>[a(z(Ne))]),_:1})]),_:2},1032,["disabled","onClick"])]),_:1})):C("",!0),j.value?(u(),R(J,{key:1,abstract:""},{default:s(({handleClick:d})=>[a(f,{disabled:x.value.length>0&&g.value==="public/video"||x.value.length===9,onClick:()=>{G("attachment"),d()},quaternary:"",circle:"",type:"primary"},{icon:s(()=>[a(r,{size:"20",color:"var(--primary-color)"},{default:s(()=>[a(z(Fe))]),_:1})]),_:2},1032,["disabled","onClick"])]),_:1})):C("",!0),a(f,{quaternary:"",circle:"",type:"primary",onClick:Y(se,["stop"])},{icon:s(()=>[a(r,{size:"20",color:"var(--primary-color)"},{default:s(()=>[a(z(Be))]),_:1})]),_:1},8,["onClick"]),H.value?(u(),R(f,{key:2,quaternary:"",circle:"",type:"primary",onClick:Y(le,["stop"])},{icon:s(()=>[a(r,{size:"20",color:"var(--primary-color)"},{default:s(()=>[a(z(Me))]),_:1})]),_:1},8,["onClick"])):C("",!0)]),T("div",vt,[a(ye,{trigger:"hover",placement:"bottom"},{trigger:s(()=>[a(ge,{class:"text-statistic",type:"circle","show-indicator":!1,status:"success","stroke-width":10,percentage:w.value.length/z(S)*100},null,8,["percentage"])]),default:s(()=>[L(" "+Q(w.value.length)+" / "+Q(z(S)),1)]),_:1}),a(f,{loading:m.value,onClick:fe,type:"primary",secondary:"",round:""},{default:s(()=>[L(" 发布 ")]),_:1},8,["loading"])])]),T("div",ft,[a(he),k.value.length>0?(u(),h("div",gt,[D.value?(u(),R(we,{key:0,value:U.value,"onUpdate:value":e[0]||(e[0]=d=>U.value=d),min:0,max:1e5,placeholder:"请输入附件价格，0为免费附件"},{prefix:s(()=>[yt]),_:1},8,["value"])):C("",!0)])):C("",!0)])]),_:1},8,["headers","data","file-list"]),v.value?(u(),h("div",ht,[a($e,{value:b.value,"onUpdate:value":e[1]||(e[1]=d=>b.value=d),name:"radiogroup"},{default:s(()=>[a(Ce,null,{default:s(()=>[(u(),h(Z,null,W(M,d=>a(ke,{key:d.value,value:d.value,label:d.label},null,8,["value","label"])),64))]),_:1})]),_:1},8,["value"])])):C("",!0),i.value?(u(),h("div",wt,[a(Te,{value:I.value,"onUpdate:value":e[2]||(e[2]=d=>I.value=d),placeholder:"请输入以http(s)://开头的链接",min:0,max:3},{"create-button-default":s(()=>[L(" 创建链接 ")]),_:1},8,["value"])])):C("",!0)])):(u(),h("div",bt,[kt,B.value?C("",!0):(u(),h("div",Ct,[a(f,{strong:"",secondary:"",round:"",type:"primary",onClick:e[3]||(e[3]=d=>K("signin"))},{default:s(()=>[L(" 登录 ")]),_:1})])),B.value?(u(),h("div",$t,[a(f,{strong:"",secondary:"",round:"",type:"primary",onClick:e[4]||(e[4]=d=>K("signin"))},{default:s(()=>[L(" 登录 ")]),_:1}),a(f,{strong:"",secondary:"",round:"",type:"info",onClick:e[5]||(e[5]=d=>K("signup"))},{default:s(()=>[L(" 注册 ")]),_:1})])):C("",!0)]))])}}});const xt={key:0,class:"skeleton-wrap"},zt={key:0,class:"empty-wrap"},It={key:1},qt={key:2},Ut={class:"load-more-wrap"},At={class:"load-more-spinner"},Et=te({__name:"Home",setup(N){const q=oe(),c=lt(),P=it(),_=l(!1),m=l(!1),i=l([]),v=l(1),w=l(20),I=l(0),F=qe(()=>{let p="泡泡广场";return c.query&&c.query.q&&(c.query.t&&c.query.t==="tag"?p="#"+decodeURIComponent(c.query.q):p="搜索: "+decodeURIComponent(c.query.q)),p}),U=()=>{_.value=!0,Pe({query:c.query.q?decodeURIComponent(c.query.q):null,type:c.query.t,page:v.value,page_size:w.value}).then(p=>{_.value=!1,p.list.length===0&&(m.value=!0),v.value>1?i.value=i.value.concat(p.list):(i.value=p.list,window.scrollTo(0,0)),I.value=Math.ceil(p.pager.total_rows/w.value)}).catch(p=>{_.value=!1,v.value>1&&v.value--})},g=p=>{if(v.value!=1){P.push({name:"post",query:{id:p.id}});return}let y=[],k=i.value.length;k==w.value&&k--;for(var b=0;b<k;b++){let $=i.value[b];if(!$.is_top)break;y.push($)}for(y.push(p);b<k;b++)y.push(i.value[b]);i.value=y},x=()=>{v.value<I.value||I.value==0?(m.value=!1,v.value++,U()):m.value=!0};return ae(()=>{U()}),Ue(()=>({path:c.path,query:c.query,refresh:q.state.refresh}),(p,y)=>{if(p.refresh!==y.refresh){m.value=!1,v.value=1,setTimeout(()=>{U()},0);return}y.path!=="/post"&&p.path==="/"&&(m.value=!1,v.value=1,setTimeout(()=>{U()},0))}),(p,y)=>{const k=st,b=Tt,$=at,M=Ie,S=ot,B=xe,O=ze,j=tt,D=nt,H=ne;return u(),h("div",null,[a(k,{title:F.value},null,8,["title"]),a(j,{class:"main-content-wrap",bordered:""},{default:s(()=>[a($,null,{default:s(()=>[a(b,{onPostSuccess:g})]),_:1}),_.value&&i.value.length===0?(u(),h("div",xt,[a(M,{num:w.value},null,8,["num"])])):C("",!0),T("div",null,[i.value.length===0?(u(),h("div",zt,[a(S,{size:"large",description:"暂无数据"})])):C("",!0),z(q).state.desktopModelShow?(u(),h("div",It,[(u(!0),h(Z,null,W(i.value,A=>(u(),R($,{key:A.id},{default:s(()=>[a(B,{post:A},null,8,["post"])]),_:2},1024))),128))])):(u(),h("div",qt,[(u(!0),h(Z,null,W(i.value,A=>(u(),R($,{key:A.id},{default:s(()=>[a(O,{post:A},null,8,["post"])]),_:2},1024))),128))]))])]),_:1}),I.value>0?(u(),R(H,{key:0,justify:"center"},{default:s(()=>[a(z(rt),{class:"load-more",slots:{complete:"没有更多泡泡了",error:"加载出错"},onInfinite:y[0]||(y[0]=A=>x())},{spinner:s(()=>[T("div",Ut,[m.value?C("",!0):(u(),R(D,{key:0,size:14})),T("span",At,Q(m.value?"没有更多泡泡了":"加载更多"),1)])]),_:1})]),_:1})):C("",!0)])}}});const da=Se(Et,[["__scopeId","data-v-8f151fd6"]]);export{da as default};