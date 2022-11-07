import{bf as oe,s as F,r as I,ar as ie,e as se,d as $,i as ae,au as G,h as R,bg as le,u as ue,j as q,A as de,t as pe,b3 as ce,b4 as ve,B as fe,C as me,bh as ye,b2 as he,V as s,W as c,Y as S,ah as H,aa as h,ab as k,a3 as i,a4 as p,a2 as m,a9 as N,a7 as x,ae as Y,a5 as y,aY as M,aZ as O,bi as X,a6 as g,bj as Q,bk as ge,bl as we,bm as be,a8 as ke,bn as xe,bo as $e,J as Ce,bp as Se}from"./index.f480f018.js";function Ee(e){if(typeof e=="number")return{"":e.toString()};const n={};return e.split(/ +/).forEach(l=>{if(l==="")return;const[a,u]=l.split(":");u===void 0?n[""]=a:n[a]=u}),n}function D(e,n){var l;if(e==null)return;const a=Ee(e);if(n===void 0)return a[""];if(typeof n=="string")return(l=a[n])!==null&&l!==void 0?l:a[""];if(Array.isArray(n)){for(let u=n.length-1;u>=0;--u){const r=n[u];if(r in a)return a[r]}return a[""]}else{let u,r=-1;return Object.keys(a).forEach(t=>{const d=Number(t);!Number.isNaN(d)&&n>=d&&d>=r&&(r=d,u=a[t])}),u}}const _e={xs:0,s:640,m:1024,l:1280,xl:1536,"2xl":1920};function Be(e){return`(min-width: ${e}px)`}const T={};function ze(e=_e){if(!oe)return F(()=>[]);if(typeof window.matchMedia!="function")return F(()=>[]);const n=I({}),l=Object.keys(e),a=(u,r)=>{u.matches?n.value[r]=!0:n.value[r]=!1};return l.forEach(u=>{const r=e[u];let t,d;T[r]===void 0?(t=window.matchMedia(Be(r)),t.addEventListener?t.addEventListener("change",v=>{d.forEach(w=>{w(v,u)})}):t.addListener&&t.addListener(v=>{d.forEach(w=>{w(v,u)})}),d=new Set,T[r]={mql:t,cbs:d}):(t=T[r].mql,d=T[r].cbs),d.add(a),t.matches&&d.forEach(v=>{v(t,u)})}),ie(()=>{l.forEach(u=>{const{cbs:r}=T[e[u]];r.has(a)&&r.delete(a)})}),F(()=>{const{value:u}=n;return l.filter(r=>u[r])})}const L=1,W=se("n-grid"),Z=1,je={span:{type:[Number,String],default:Z},offset:{type:[Number,String],default:0},suffix:Boolean,privateOffset:Number,privateSpan:Number,privateColStart:Number,privateShow:{type:Boolean,default:!0}};var J=$({__GRID_ITEM__:!0,name:"GridItem",alias:["Gi"],props:je,setup(){const{xGapRef:e,itemStyleRef:n,overflowRef:l}=ae(W),a=le();return{overflow:l,itemStyle:n,deriveStyle:()=>{const{privateSpan:u=Z,privateShow:r=!0,privateColStart:t=void 0,privateOffset:d=0}=a.vnode.props,{value:v}=e,w=G(v||0);return{display:r?"":"none",gridColumn:`${t!=null?t:`span ${u}`} / span ${u}`,marginLeft:d?`calc((100% - (${u} - 1) * ${w}) / ${u} * ${d} + ${w} * ${d})`:""}}}},render(){var e,n;return R("div",{style:[this.itemStyle,this.deriveStyle()]},(n=(e=this.$slots).default)===null||n===void 0?void 0:n.call(e,{overflow:this.overflow}))}});const Re={xs:0,s:640,m:1024,l:1280,xl:1536,xxl:1920},K=24,Fe={responsive:{type:[String,Boolean],default:"self"},cols:{type:[Number,String],default:K},itemResponsive:Boolean,collapsed:Boolean,collapsedRows:{type:Number,default:1},itemStyle:[Object,String],xGap:{type:[Number,String],default:0},yGap:{type:[Number,String],default:0}};var ee=$({name:"Grid",inheritAttrs:!1,props:Fe,setup(e){const{mergedClsPrefixRef:n,mergedBreakpointsRef:l}=ue(e),a=/^\d+$/,u=I(void 0),r=ze((l==null?void 0:l.value)||Re),t=q(()=>!!(e.itemResponsive||!a.test(e.cols.toString())||!a.test(e.xGap.toString())||!a.test(e.yGap.toString()))),d=F(()=>{if(!!t.value)return e.responsive==="self"?u.value:r.value}),v=q(()=>{var C;return(C=Number(D(e.cols.toString(),d.value)))!==null&&C!==void 0?C:K}),w=q(()=>D(e.xGap.toString(),d.value)),o=q(()=>D(e.yGap.toString(),d.value)),f=C=>{u.value=C.contentRect.width},E=C=>{ve(f,C)},z=I(!1),_=F(()=>{if(e.responsive==="self")return E});return de(W,{itemStyleRef:pe(e,"itemStyle"),xGapRef:w,overflowRef:z}),{mergedClsPrefix:n,style:F(()=>({width:"100%",display:"grid",gridTemplateColumns:`repeat(${v.value}, minmax(0, 1fr))`,columnGap:G(w.value),rowGap:G(o.value)})),isResponsive:t,responsiveQuery:d,responsiveCols:v,handleResize:_,overflow:z}},render(){const e=()=>{var n,l,a,u,r,t;this.overflow=!1;const d=fe(me(this)),v=[],{collapsed:w,collapsedRows:o,responsiveCols:f,responsiveQuery:E}=this;d.forEach(b=>{var A,j,B;if(((A=b==null?void 0:b.type)===null||A===void 0?void 0:A.__GRID_ITEM__)!==!0)return;const P=ye(b),U=Number((B=D((j=P.props)===null||j===void 0?void 0:j.span,E))!==null&&B!==void 0?B:L);U!==0&&v.push({child:P,rawChildSpan:U})});let z=0;const _=(n=v[v.length-1])===null||n===void 0?void 0:n.child;if(_!=null&&_.props){const b=(l=_.props)===null||l===void 0?void 0:l.suffix;b!==void 0&&b!==!1&&(z=(u=(a=_.props)===null||a===void 0?void 0:a.span)!==null&&u!==void 0?u:L,_.props.privateSpan=z,_.props.privateColStart=f+1-z,_.props.privateShow=!0)}let C=0,V=!1;for(const{child:b,rawChildSpan:A}of v){if(V&&(this.overflow=!0),!V){const j=Number((t=D((r=b.props)===null||r===void 0?void 0:r.offset,E))!==null&&t!==void 0?t:0),B=Math.min(A+j,f)||1;if(b.props?(b.props.privateSpan=B,b.props.privateOffset=j):b.props={privateSpan:B,privateOffset:j},w){const P=C%f;B+P>f&&(C+=f-P),B+C+z>o*f?V=!0:C+=B}}V&&(b.props?b.props.privateShow!==!0&&(b.props.privateShow=!1):b.props={privateShow:!1})}return R("div",he({class:`${this.mergedClsPrefix}-grid`,style:this.style},this.$attrs),v.map(({child:b})=>b))};return this.isResponsive&&this.responsive==="self"?R(ce,{onResize:this.handleResize},{default:e}):e()}});const Ie={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Me=S("path",{d:"M352 48H160a48 48 0 0 0-48 48v368l144-128l144 128V96a48 48 0 0 0-48-48z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),Ne=[Me];var jn=$({name:"BookmarkOutline",render:function(n,l){return s(),c("svg",Ie,Ne)}});const Ae={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Pe=S("path",{d:"M408 64H104a56.16 56.16 0 0 0-56 56v192a56.16 56.16 0 0 0 56 56h40v80l93.72-78.14a8 8 0 0 1 5.13-1.86H408a56.16 56.16 0 0 0 56-56V120a56.16 56.16 0 0 0-56-56z",fill:"none",stroke:"currentColor","stroke-linejoin":"round","stroke-width":"32"},null,-1),De=[Pe];var Rn=$({name:"ChatboxOutline",render:function(n,l){return s(),c("svg",Ae,De)}});const Te={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Ve=S("path",{d:"M320 336h76c55 0 100-21.21 100-75.6s-53-73.47-96-75.6C391.11 99.74 329 48 256 48c-69 0-113.44 45.79-128 91.2c-60 5.7-112 35.88-112 98.4S70 336 136 336h56",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),qe=S("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M192 400.1l64 63.9l64-63.9"},null,-1),Oe=S("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M256 224v224.03"},null,-1),Ge=[Ve,qe,Oe];var Ue=$({name:"CloudDownloadOutline",render:function(n,l){return s(),c("svg",Te,Ge)}});const Le={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},He=S("path",{d:"M352.92 80C288 80 256 144 256 144s-32-64-96.92-64c-52.76 0-94.54 44.14-95.08 96.81c-1.1 109.33 86.73 187.08 183 252.42a16 16 0 0 0 18 0c96.26-65.34 184.09-143.09 183-252.42c-.54-52.67-42.32-96.81-95.08-96.81z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),Ye=[He];var Fn=$({name:"HeartOutline",render:function(n,l){return s(),c("svg",Le,Ye)}});const Xe={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Qe=S("path",{d:"M208 352h-64a96 96 0 0 1 0-192h64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),We=S("path",{d:"M304 160h64a96 96 0 0 1 0 192h-64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),Ze=S("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36",d:"M163.29 256h187.42"},null,-1),Je=[Qe,We,Ze];var Ke=$({name:"LinkOutline",render:function(n,l){return s(),c("svg",Xe,Je)}});const en={class:"link-wrap"},nn=["href"],tn={class:"link-txt"},rn=$({props:{links:{default:()=>[]}},setup(e){const n=e;return(l,a)=>{const u=Y;return s(),c("div",en,[(s(!0),c(h,null,k(n.links,r=>(s(),c("div",{class:"link-item",key:r.id},[i(u,{class:"hash-link"},{default:p(()=>[i(m(Ke))]),_:1}),S("a",{href:r.content,class:"hash-link",target:"_blank",onClick:a[0]||(a[0]=x(()=>{},["stop"]))},[S("span",tn,N(r.content),1)],8,nn)]))),128))])}}});var In=H(rn,[["__scopeId","data-v-4c9a59cc"]]),ne=$({name:"BasicTheme",props:{uuid:{type:String,required:!0},src:{type:String,required:!0},autoplay:{type:Boolean,required:!0},loop:{type:Boolean,required:!0},controls:{type:Boolean,required:!0},hoverable:{type:Boolean,required:!0},mask:{type:Boolean,required:!0},colors:{type:[String,Array],required:!0},time:{type:Object,required:!0},playing:{type:Boolean,default:!1},duration:{type:[String,Number],required:!0}},data(){return{hovered:!1,volume:!1,amount:1}},computed:{colorFrom(){var e;return typeof this.colors=="string"?this.colors?this.colors:"#fbbf24":(e=this.colors)!==null&&e!==void 0&&e[0]?this.colors[0]:"#fbbf24"},colorTo(){var e;return typeof this.colors=="string"?this.colors?this.colors:"#fbbf24":(e=this.colors)!==null&&e!==void 0&&e[1]?this.colors[1]:"#ec4899"}},mounted(){this.$emit("setPlayer",this.$refs[this.uuid])},methods:{setVolume(){this.$refs[this.uuid].volume=this.amount},stopVolume(){return this.amount>0?this.amount=0:this.amount=1}}});const on={class:"relative"},sn={class:"flex items-center justify-start w-full"},an={class:"font-sans text-white text-xs w-24"},ln={class:"mr-3 ml-2"},un={class:"relative"},dn={class:"px-3 py-2 rounded-lg flex items-center transform translate-x-2",style:{"background-color":"rgba(0, 0, 0, .8)"}},pn=i("img",{src:"https://en-zo.dev/vue-videoplayer/play.svg",alt:"Icon play video",class:"transform translate-x-0.5 w-12"},null,-1);function cn(e,n,l,a,u,r){return s(),y("div",{class:"shadow-xl rounded-xl overflow-hidden relative",onMouseenter:n[15]||(n[15]=t=>e.hovered=!0),onMouseleave:n[16]||(n[16]=t=>e.hovered=!1),onKeydown:n[17]||(n[17]=Q(t=>e.$emit("play"),["left"]))},[i("div",on,[i("video",{ref:e.uuid,class:"w-full",loop:e.loop,autoplay:e.autoplay,muted:e.autoplay,onTimeupdate:n[1]||(n[1]=t=>e.$emit("timeupdate",t.target)),onPause:n[2]||(n[2]=t=>e.$emit("isPlaying",!1)),onPlay:n[3]||(n[3]=t=>e.$emit("isPlaying",!0)),onClick:n[4]||(n[4]=t=>e.$emit("play"))},[i("source",{src:e.src,type:"video/mp4"},null,8,["src"])],40,["loop","autoplay","muted"]),e.controls?(s(),y("div",{key:0,class:[{"opacity-0 translate-y-full":!e.hoverable&&e.hovered,"opacity-0 translate-y-full":e.hoverable&&!e.hovered},"transition duration-300 transform absolute w-full bottom-0 left-0 flex items-center justify-between overlay px-5  pt-3 pb-5"]},[i("div",sn,[i("p",an,N(e.time.display)+"/"+N(e.duration),1),i("div",ln,[M(i("img",{src:"https://en-zo.dev/vue-videoplayer/pause.svg",alt:"Icon pause video",class:"w-5 cursor-pointer",onClick:n[5]||(n[5]=t=>e.$emit("play"))},null,512),[[O,e.playing]]),M(i("img",{src:"https://en-zo.dev/vue-videoplayer/play.svg",alt:"Icon play video",class:"w-5 cursor-pointer",onClick:n[6]||(n[6]=t=>e.$emit("play"))},null,512),[[O,!e.playing]])]),i("div",{class:"w-full h-1 bg-white bg-opacity-60 rounded-sm cursor-pointer",onClick:n[7]||(n[7]=t=>e.$emit("position",t))},[i("div",{class:"relative h-full pointer-events-none",style:`width: ${e.time.progress}%; transition: width .2s ease-in-out;`},[i("div",{class:"w-full rounded-sm h-full gradient-variable bg-gradient-to-r pointer-events-none absolute top-0 left-0",style:`--tw-gradient-from: ${e.colorFrom};--tw-gradient-to: ${e.colorTo};transition: width .2s ease-in-out`},null,4),i("div",{class:"w-full rounded-sm filter blur-sm h-full gradient-variable bg-gradient-to-r pointer-events-none absolute top-0 left-0",style:`--tw-gradient-from: ${e.colorFrom};--tw-gradient-to: ${e.colorTo};transition: width .2s ease-in-out`},null,4)],4)])]),i("div",{class:"ml-5 flex items-center justify-end",onMouseleave:n[13]||(n[13]=t=>e.volume=!1)},[i("div",un,[i("div",{class:`w-128 origin-left translate-x-2 -rotate-90 w-128 transition duration-200 absolute transform top-0 py-2 ${e.volume?"-translate-y-4":"opacity-0 -translate-y-1 pointer-events-none"}`},[i("div",dn,[M(i("input",{"onUpdate:modelValue":n[8]||(n[8]=t=>e.amount=t),type:"range",step:"0.05",min:"0",max:"1",class:"rounded-lg overflow-hidden appearance-none bg-white bg-opacity-30 h-1 w-128 vertical-range",onInput:n[9]||(n[9]=(...t)=>e.setVolume&&e.setVolume(...t))},null,544),[[X,e.amount]])])],2),i("img",{src:`https://en-zo.dev/vue-videoplayer/volume-${Math.ceil(e.amount*2)}.svg`,alt:"High volume video",class:"w-5 cursor-pointer relative",style:{"z-index":"2"},onClick:n[10]||(n[10]=(...t)=>e.stopVolume&&e.stopVolume(...t)),onMouseenter:n[11]||(n[11]=t=>e.volume=!0)},null,40,["src"])]),i("img",{src:"https://en-zo.dev/vue-videoplayer/maximize.svg",alt:"Fullscreen",class:"w-3 ml-4 cursor-pointer",onClick:n[12]||(n[12]=t=>e.$emit("fullScreen"))})],32)],2)):g("",!0),!e.autoplay&&e.mask&&e.time.current===0?(s(),y("div",{key:1,class:`transition transform duration-300 absolute top-0 left-0 w-full h-full bg-black bg-opacity-50 backdrop-filter z-10 flex items-center justify-center ${e.playing?"opacity-0 pointer-events-none":""}`},[i("div",{class:"w-20 h-20 rounded-full bg-white bg-opacity-20 transition duration-200 hover:bg-opacity-40 flex items-center justify-center cursor-pointer",onClick:n[14]||(n[14]=t=>e.$emit("play"))},[pn])],2)):g("",!0)])],32)}ne.render=cn;var te=$({name:"BasicTheme",props:{uuid:{type:String,required:!0},src:{type:String,required:!0},autoplay:{type:Boolean,required:!0},loop:{type:Boolean,required:!0},controls:{type:Boolean,required:!0},hoverable:{type:Boolean,required:!0},mask:{type:Boolean,required:!0},colors:{type:[String,Array],required:!0},time:{type:Object,required:!0},playing:{type:Boolean,default:!1},duration:{type:[String,Number],required:!0}},data(){return{hovered:!1,volume:!1,amount:1}},computed:{color(){var e;return typeof this.colors=="string"?this.colors?this.colors:"#8B5CF6":(e=this.colors)!==null&&e!==void 0&&e[0]?this.colors[0]:"#8B5CF6"}},mounted(){this.$emit("setPlayer",this.$refs[this.uuid])},methods:{setVolume(){this.$refs[this.uuid].volume=this.amount},stopVolume(){return this.amount>0?this.amount=0:this.amount=1}}});const vn={class:"relative"},fn={class:"mr-5"},mn={class:"relative mr-6"},yn={class:"px-3 py-3 rounded-xl flex items-center transform translate-x-9 bg-black bg-opacity-30"},hn=i("img",{src:"https://en-zo.dev/vue-videoplayer/play.svg",alt:"Icon play video",class:"transform translate-x-0.5 w-12"},null,-1);function gn(e,n,l,a,u,r){return s(),y("div",{class:"shadow-xl rounded-3xl overflow-hidden relative",onMouseenter:n[14]||(n[14]=t=>e.hovered=!0),onMouseleave:n[15]||(n[15]=t=>e.hovered=!1),onKeydown:n[16]||(n[16]=Q(t=>e.$emit("play"),["left"]))},[i("div",vn,[i("video",{ref:e.uuid,class:"w-full",loop:e.loop,autoplay:e.autoplay,muted:e.autoplay,onTimeupdate:n[1]||(n[1]=t=>e.$emit("timeupdate",t.target)),onPause:n[2]||(n[2]=t=>e.$emit("isPlaying",!1)),onPlay:n[3]||(n[3]=t=>e.$emit("isPlaying",!0)),onClick:n[4]||(n[4]=t=>e.$emit("play"))},[i("source",{src:e.src,type:"video/mp4"},null,8,["src"])],40,["loop","autoplay","muted"]),e.controls?(s(),y("div",{key:0,class:[{"opacity-0 translate-y-full":!e.hoverable&&e.hovered,"opacity-0 translate-y-full":e.hoverable&&!e.hovered},"absolute px-5 pb-5 bottom-0 left-0 w-full transition duration-300 transform"]},[i("div",{class:"w-full bg-black bg-opacity-30 px-5 py-4 rounded-xl flex items-center justify-between",onMouseleave:n[12]||(n[12]=t=>e.volume=!1)},[i("div",{class:"font-sans py-1 px-2 text-white rounded-md text-xs mr-5 whitespace-nowrap font-medium w-32 text-center",style:`font-size: 11px; background-color: ${e.color}`},N(e.time.display)+"\xA0/\xA0"+N(e.duration),5),i("div",fn,[M(i("img",{src:"https://en-zo.dev/vue-videoplayer/basic/pause.svg",alt:"Icon pause video",class:"w-4 cursor-pointer filter-white transition duration-300",onClick:n[5]||(n[5]=t=>e.$emit("play"))},null,512),[[O,e.playing]]),M(i("img",{src:"https://en-zo.dev/vue-videoplayer/basic/play.svg",alt:"Icon play video",class:"w-4 cursor-pointer filter-white transition duration-300",onClick:n[6]||(n[6]=t=>e.$emit("play"))},null,512),[[O,!e.playing]])]),i("div",{class:"w-full h-1 bg-white bg-opacity-40 rounded-sm cursor-pointer mr-6",onClick:n[7]||(n[7]=t=>e.$emit("position",t))},[i("div",{class:"w-full rounded-sm h-full bg-white pointer-events-none",style:`width: ${e.time.progress}%; transition: width .2s ease-in-out;`},null,4)]),i("div",mn,[i("div",{class:`w-128 origin-left translate-x-2 -rotate-90 w-128 transition duration-200 absolute transform top-0 py-2 ${e.volume?"-translate-y-4":"opacity-0 -translate-y-1 pointer-events-none"}`},[i("div",yn,[M(i("input",{"onUpdate:modelValue":n[8]||(n[8]=t=>e.amount=t),type:"range",step:"0.05",min:"0",max:"1",class:"rounded-lg overflow-hidden appearance-none bg-white bg-opacity-30 h-1.5 w-128 vertical-range"},null,512),[[X,e.amount]])])],2),i("img",{src:`https://en-zo.dev/vue-videoplayer/basic/volume_${Math.ceil(e.amount*2)}.svg`,alt:"High volume video",class:"w-5 cursor-pointer filter-white transition duration-300 relative",style:{"z-index":"2"},onClick:n[9]||(n[9]=(...t)=>e.stopVolume&&e.stopVolume(...t)),onMouseenter:n[10]||(n[10]=t=>e.volume=!0)},null,40,["src"])]),i("img",{src:"https://en-zo.dev/vue-videoplayer/basic/fullscreen.svg",alt:"Fullscreen",class:"w-4 cursor-pointer filter-white transition duration-300",onClick:n[11]||(n[11]=t=>e.$emit("fullScreen"))})],32)],2)):g("",!0),!e.autoplay&&e.mask&&e.time.current===0?(s(),y("div",{key:1,class:`transition transform duration-300 absolute top-0 left-0 w-full h-full bg-black bg-opacity-50 backdrop-filter z-10 flex items-center justify-center ${e.playing?"opacity-0 pointer-events-none":""}`},[i("div",{class:"w-20 h-20 rounded-full bg-white bg-opacity-20 transition duration-200 hover:bg-opacity-40 flex items-center justify-center cursor-pointer",onClick:n[13]||(n[13]=t=>e.$emit("play"))},[hn])],2)):g("",!0)])],32)}te.render=gn;var re=$({name:"Vue3PlayerVideo",components:{basic:te,gradient:ne},props:{src:{type:String,required:!0},autoplay:{type:Boolean,default:!1},loop:{type:Boolean,default:!1},controls:{type:Boolean,default:!0},mask:{type:Boolean,default:!0},colors:{type:[String,Array],default(){return["#8B5CF6","#ec4899"]}},hoverable:{type:Boolean,default:!1},theme:{type:String,default:"basic"}},data(){return{uuid:Math.random().toString(36).substr(2,18),player:null,duration:0,playing:!1,time:{progress:0,display:0,current:0}}},watch:{"time.current"(e){this.time.display=this.format(Number(e)),this.time.progress=e*100/this.player.duration}},methods:{isPlaying(e){this.playing=e},play(){return this.playing?this.player.pause():this.player.play()},setPlayer(e){this.player=e,this.player.addEventListener("loadeddata",()=>{this.player.readyState>=3&&(this.duration=this.format(Number(this.player.duration)),this.time.display=this.format(0))})},stop(){this.player.pause(),this.player.currentTime=0},fullScreen(){this.player.webkitEnterFullscreen()},position(e){this.player.pause();const n=e.target.getBoundingClientRect(),a=(e.clientX-n.left)*100/e.target.offsetWidth;this.player.currentTime=a*this.player.duration/100,this.player.play()},format(e){const n=Math.floor(e/3600),l=Math.floor(e%3600/60),a=Math.round(e%60);return[n,l>9?l:n?"0"+l:l||"00",a>9?a:"0"+a].filter(Boolean).join(":")}}});const wn={class:"vue3-player-video"};function bn(e,n,l,a,u,r){return s(),y("div",wn,[(s(),y(ge(e.theme),{uuid:e.uuid,src:e.src,autoplay:e.autoplay,loop:e.loop,controls:e.controls,mask:e.mask,colors:e.colors,time:e.time,playing:e.playing,duration:e.duration,hoverable:e.hoverable,onPlay:e.play,onStop:e.stop,onTimeupdate:n[1]||(n[1]=({currentTime:t})=>e.time.current=t),onPosition:e.position,onFullScreen:e.fullScreen,onSetPlayer:e.setPlayer,onIsPlaying:e.isPlaying},null,8,["uuid","src","autoplay","loop","controls","mask","colors","time","playing","duration","hoverable","onPlay","onStop","onPosition","onFullScreen","onSetPlayer","onIsPlaying"]))])}function kn(e,n){n===void 0&&(n={});var l=n.insertAt;if(!(!e||typeof document=="undefined")){var a=document.head||document.getElementsByTagName("head")[0],u=document.createElement("style");u.type="text/css",l==="top"&&a.firstChild?a.insertBefore(u,a.firstChild):a.appendChild(u),u.styleSheet?u.styleSheet.cssText=e:u.appendChild(document.createTextNode(e))}}var xn=`/*! tailwindcss v2.1.2 | MIT License | https://tailwindcss.com */

/*! modern-normalize v1.1.0 | MIT License | https://github.com/sindresorhus/modern-normalize */

/*
Document
========
*/

/**
Use a better box model (opinionated).
*/

*,
::before,
::after {
  box-sizing: border-box;
}

/**
Use a more readable tab size (opinionated).
*/

/**
1. Correct the line height in all browsers.
2. Prevent adjustments of font size after orientation changes in iOS.
*/

/*
Sections
========
*/

/**
Remove the margin in all browsers.
*/

/**
Improve consistency of default fonts in all browsers. (https://github.com/sindresorhus/modern-normalize/issues/3)
*/

/*
Grouping content
================
*/

/**
1. Add the correct height in Firefox.
2. Correct the inheritance of border color in Firefox. (https://bugzilla.mozilla.org/show_bug.cgi?id=190655)
*/

/*
Text-level semantics
====================
*/

/**
Add the correct text decoration in Chrome, Edge, and Safari.
*/

/**
Add the correct font weight in Edge and Safari.
*/

/**
1. Improve consistency of default fonts in all browsers. (https://github.com/sindresorhus/modern-normalize/issues/3)
2. Correct the odd 'em' font sizing in all browsers.
*/

/**
Add the correct font size in all browsers.
*/

/**
Prevent 'sub' and 'sup' elements from affecting the line height in all browsers.
*/

/*
Tabular data
============
*/

/**
1. Remove text indentation from table contents in Chrome and Safari. (https://bugs.chromium.org/p/chromium/issues/detail?id=999088, https://bugs.webkit.org/show_bug.cgi?id=201297)
2. Correct table border color inheritance in all Chrome and Safari. (https://bugs.chromium.org/p/chromium/issues/detail?id=935729, https://bugs.webkit.org/show_bug.cgi?id=195016)
*/

/*
Forms
=====
*/

/**
1. Change the font styles in all browsers.
2. Remove the margin in Firefox and Safari.
*/


input {
  font-family: inherit; /* 1 */
  font-size: 100%; /* 1 */
  line-height: 1.15; /* 1 */
  margin: 0; /* 2 */
}

/**
Remove the inheritance of text transform in Edge and Firefox.
1. Remove the inheritance of text transform in Firefox.
*/

/**
Correct the inability to style clickable types in iOS and Safari.
*/


[type='button'],
[type='reset'],
[type='submit'] {
  -webkit-appearance: button;
}

/**
Remove the inner border and padding in Firefox.
*/

::-moz-focus-inner {
  border-style: none;
  padding: 0;
}

/**
Restore the focus styles unset by the previous rule.
*/

:-moz-focusring {
  outline: 1px dotted ButtonText;
}

/**
Remove the additional ':invalid' styles in Firefox.
See: https://github.com/mozilla/gecko-dev/blob/2f9eacd9d3d995c937b4251a5557d95d494c9be1/layout/style/res/forms.css#L728-L737
*/

:-moz-ui-invalid {
  box-shadow: none;
}

/**
Remove the padding so developers are not caught out when they zero out 'fieldset' elements in all browsers.
*/

/**
Add the correct vertical alignment in Chrome and Firefox.
*/

progress {
  vertical-align: baseline;
}

/**
Correct the cursor style of increment and decrement buttons in Safari.
*/

::-webkit-inner-spin-button,
::-webkit-outer-spin-button {
  height: auto;
}

/**
1. Correct the odd appearance in Chrome and Safari.
2. Correct the outline style in Safari.
*/

[type='search'] {
  -webkit-appearance: textfield; /* 1 */
  outline-offset: -2px; /* 2 */
}

/**
Remove the inner padding in Chrome and Safari on macOS.
*/

::-webkit-search-decoration {
  -webkit-appearance: none;
}

/**
1. Correct the inability to style clickable types in iOS and Safari.
2. Change font properties to 'inherit' in Safari.
*/

::-webkit-file-upload-button {
  -webkit-appearance: button; /* 1 */
  font: inherit; /* 2 */
}

/*
Interactive
===========
*/

/*
Add the correct display in Chrome and Safari.
*/

/**
 * Manually forked from SUIT CSS Base: https://github.com/suitcss/base
 * A thin layer on top of normalize.css that provides a starting point more
 * suitable for web applications.
 */

/**
 * Removes the default spacing and border for appropriate elements.
 */


p {
  margin: 0;
}

/**
 * Work around a Firefox/IE bug where the transparent \`button\` background
 * results in a loss of the default \`button\` focus styles.
 */

/**
 * Tailwind custom reset styles
 */

/**
 * 1. Use the user's configured \`sans\` font-family (with Tailwind's default
 *    sans-serif font stack as a fallback) as a sane default.
 * 2. Use Tailwind's default "normal" line-height so the user isn't forced
 *    to override it to ensure consistency even when using the default theme.
 */

/**
 * Inherit font-family and line-height from \`html\` so users can set them as
 * a class directly on the \`html\` element.
 */

/**
 * 1. Prevent padding and border from affecting element width.
 *
 *    We used to set this in the html element and inherit from
 *    the parent element for everything else. This caused issues
 *    in shadow-dom-enhanced elements like <details> where the content
 *    is wrapped by a div with box-sizing set to \`content-box\`.
 *
 *    https://github.com/mozdevs/cssremedy/issues/4
 *
 *
 * 2. Allow adding a border to an element by just adding a border-width.
 *
 *    By default, the way the browser specifies that an element should have no
 *    border is by setting it's border-style to \`none\` in the user-agent
 *    stylesheet.
 *
 *    In order to easily add borders to elements by just setting the \`border-width\`
 *    property, we change the default border-style for all elements to \`solid\`, and
 *    use border-width to hide them instead. This way our \`border\` utilities only
 *    need to set the \`border-width\` property instead of the entire \`border\`
 *    shorthand, making our border utilities much more straightforward to compose.
 *
 *    https://github.com/tailwindcss/tailwindcss/pull/116
 */

*,
::before,
::after {
  box-sizing: border-box; /* 1 */
  border-width: 0; /* 2 */
  border-style: solid; /* 2 */
  border-color: #e5e7eb; /* 2 */
}

/*
 * Ensure horizontal rules are visible by default
 */

/**
 * Undo the \`border-style: none\` reset that Normalize applies to images so that
 * our \`border-{width}\` utilities have the expected effect.
 *
 * The Normalize reset is unnecessary for us since we default the border-width
 * to 0 on all elements.
 *
 * https://github.com/tailwindcss/tailwindcss/issues/362
 */

img {
  border-style: solid;
}

input:-ms-input-placeholder {
  opacity: 1;
  color: #9ca3af;
}

input::placeholder {
  opacity: 1;
  color: #9ca3af;
}


[role="button"] {
  cursor: pointer;
}

/**
 * Reset links to optimize for opt-in styling instead of
 * opt-out.
 */

/**
 * Reset form element properties that are easy to forget to
 * style explicitly so you don't inadvertently introduce
 * styles that deviate from your design system. These styles
 * supplement a partial reset that is already applied by
 * normalize.css.
 */


input {
  padding: 0;
  line-height: inherit;
  color: inherit;
}

/**
 * Use the configured 'mono' font family for elements that
 * are expected to be rendered with a monospace font, falling
 * back to the system monospace stack if there is no configured
 * 'mono' font family.
 */

/**
 * Make replaced elements \`display: block\` by default as that's
 * the behavior you want almost all of the time. Inspired by
 * CSS Remedy, with \`svg\` added as well.
 *
 * https://github.com/mozdevs/cssremedy/issues/14
 */

img,
svg,
video {
  display: block;
  vertical-align: middle;
}

/**
 * Constrain images and videos to the parent width and preserve
 * their intrinsic aspect ratio.
 *
 * https://github.com/mozdevs/cssremedy/issues/14
 */

img,
video {
  max-width: 100%;
  height: auto;
}

.vue3-player-video .appearance-none{
  -webkit-appearance: none;
          appearance: none;
}

.vue3-player-video .bg-black{
  --tw-bg-opacity: 1;
  background-color: rgba(0, 0, 0, var(--tw-bg-opacity));
}

.vue3-player-video .bg-white{
  --tw-bg-opacity: 1;
  background-color: rgba(255, 255, 255, var(--tw-bg-opacity));
}

.vue3-player-video .bg-gradient-to-r{
  background-image: linear-gradient(to right, var(--tw-gradient-stops));
}

.vue3-player-video .bg-opacity-20{
  --tw-bg-opacity: 0.2;
}

.vue3-player-video .bg-opacity-30{
  --tw-bg-opacity: 0.3;
}

.vue3-player-video .bg-opacity-40{
  --tw-bg-opacity: 0.4;
}

.vue3-player-video .bg-opacity-50{
  --tw-bg-opacity: 0.5;
}

.vue3-player-video .bg-opacity-60{
  --tw-bg-opacity: 0.6;
}

.vue3-player-video .hover\\:bg-opacity-40:hover{
  --tw-bg-opacity: 0.4;
}

.vue3-player-video .rounded-sm{
  border-radius: 0.125rem;
}

.vue3-player-video .rounded-md{
  border-radius: 0.375rem;
}

.vue3-player-video .rounded-lg{
  border-radius: 0.5rem;
}

.vue3-player-video .rounded-xl{
  border-radius: 0.75rem;
}

.vue3-player-video .rounded-3xl{
  border-radius: 1.5rem;
}

.vue3-player-video .rounded-full{
  border-radius: 9999px;
}

.vue3-player-video .cursor-pointer{
  cursor: pointer;
}

.vue3-player-video .flex{
  display: flex;
}

.vue3-player-video .items-center{
  align-items: center;
}

.vue3-player-video .justify-start{
  justify-content: flex-start;
}

.vue3-player-video .justify-end{
  justify-content: flex-end;
}

.vue3-player-video .justify-center{
  justify-content: center;
}

.vue3-player-video .justify-between{
  justify-content: space-between;
}

.vue3-player-video .font-sans{
  font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
}

.vue3-player-video .font-medium{
  font-weight: 500;
}

.vue3-player-video .h-1{
  height: 0.25rem;
}

.vue3-player-video .h-20{
  height: 5rem;
}

.vue3-player-video .h-full{
  height: 100%;
}

.vue3-player-video .text-xs{
  font-size: 0.75rem;
  line-height: 1rem;
}

.vue3-player-video .ml-2{
  margin-left: 0.5rem;
}

.vue3-player-video .mr-3{
  margin-right: 0.75rem;
}

.vue3-player-video .ml-4{
  margin-left: 1rem;
}

.vue3-player-video .mr-5{
  margin-right: 1.25rem;
}

.vue3-player-video .ml-5{
  margin-left: 1.25rem;
}

.vue3-player-video .mr-6{
  margin-right: 1.5rem;
}

.vue3-player-video .opacity-0{
  opacity: 0;
}

.vue3-player-video .overflow-hidden{
  overflow: hidden;
}

.vue3-player-video .py-1{
  padding-top: 0.25rem;
  padding-bottom: 0.25rem;
}

.vue3-player-video .py-2{
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
}

.vue3-player-video .px-2{
  padding-left: 0.5rem;
  padding-right: 0.5rem;
}

.vue3-player-video .py-3{
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
}

.vue3-player-video .px-3{
  padding-left: 0.75rem;
  padding-right: 0.75rem;
}

.vue3-player-video .py-4{
  padding-top: 1rem;
  padding-bottom: 1rem;
}

.vue3-player-video .px-5{
  padding-left: 1.25rem;
  padding-right: 1.25rem;
}

.vue3-player-video .pt-3{
  padding-top: 0.75rem;
}

.vue3-player-video .pb-5{
  padding-bottom: 1.25rem;
}

.vue3-player-video .pointer-events-none{
  pointer-events: none;
}

.vue3-player-video .absolute{
  position: absolute;
}

.vue3-player-video .relative{
  position: relative;
}

.vue3-player-video .top-0{
  top: 0px;
}

.vue3-player-video .bottom-0{
  bottom: 0px;
}

.vue3-player-video .left-0{
  left: 0px;
}

*{
  --tw-shadow: 0 0 #0000;
}

.vue3-player-video .shadow-xl{
  --tw-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
}

*{
  --tw-ring-inset: var(--tw-empty,/*!*/ /*!*/);
  --tw-ring-offset-width: 0px;
  --tw-ring-offset-color: #fff;
  --tw-ring-color: rgba(59, 130, 246, 0.5);
  --tw-ring-offset-shadow: 0 0 #0000;
  --tw-ring-shadow: 0 0 #0000;
}

.vue3-player-video .text-center{
  text-align: center;
}

.vue3-player-video .text-white{
  --tw-text-opacity: 1;
  color: rgba(255, 255, 255, var(--tw-text-opacity));
}

.vue3-player-video .whitespace-nowrap{
  white-space: nowrap;
}

.vue3-player-video .w-3{
  width: 0.75rem;
}

.vue3-player-video .w-4{
  width: 1rem;
}

.vue3-player-video .w-5{
  width: 1.25rem;
}

.vue3-player-video .w-12{
  width: 3rem;
}

.vue3-player-video .w-20{
  width: 5rem;
}

.vue3-player-video .w-24{
  width: 6rem;
}

.vue3-player-video .w-32{
  width: 8rem;
}

.vue3-player-video .w-full{
  width: 100%;
}

.vue3-player-video .z-10{
  z-index: 10;
}

.vue3-player-video .transform{
  --tw-translate-x: 0;
  --tw-translate-y: 0;
  --tw-rotate: 0;
  --tw-skew-x: 0;
  --tw-skew-y: 0;
  --tw-scale-x: 1;
  --tw-scale-y: 1;
  transform: translateX(var(--tw-translate-x)) translateY(var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y));
}

.vue3-player-video .origin-left{
  transform-origin: left;
}

.vue3-player-video .-rotate-90{
  --tw-rotate: -90deg;
}

.vue3-player-video .translate-x-0{
  --tw-translate-x: 0px;
}

.vue3-player-video .translate-x-2{
  --tw-translate-x: 0.5rem;
}

.vue3-player-video .translate-x-9{
  --tw-translate-x: 2.25rem;
}

.vue3-player-video .-translate-y-1{
  --tw-translate-y: -0.25rem;
}

.vue3-player-video .-translate-y-4{
  --tw-translate-y: -1rem;
}

.vue3-player-video .translate-y-full{
  --tw-translate-y: 100%;
}

.vue3-player-video .transition{
  transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, -webkit-backdrop-filter;
  transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter, -webkit-backdrop-filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

.vue3-player-video .ease-in-out{
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

.vue3-player-video .duration-200{
  transition-duration: 200ms;
}

.vue3-player-video .duration-300{
  transition-duration: 300ms;
}

@keyframes spin{
  to{
    transform: rotate(360deg);
  }
}

@keyframes ping{
  75%, 100%{
    transform: scale(2);
    opacity: 0;
  }
}

@keyframes pulse{
  50%{
    opacity: .5;
  }
}

@keyframes bounce{
  0%, 100%{
    transform: translateY(-25%);
    animation-timing-function: cubic-bezier(0.8,0,1,1);
  }

  50%{
    transform: none;
    animation-timing-function: cubic-bezier(0,0,0.2,1);
  }
}

.vue3-player-video .filter{
  --tw-blur: var(--tw-empty,/*!*/ /*!*/);
  --tw-brightness: var(--tw-empty,/*!*/ /*!*/);
  --tw-contrast: var(--tw-empty,/*!*/ /*!*/);
  --tw-grayscale: var(--tw-empty,/*!*/ /*!*/);
  --tw-hue-rotate: var(--tw-empty,/*!*/ /*!*/);
  --tw-invert: var(--tw-empty,/*!*/ /*!*/);
  --tw-saturate: var(--tw-empty,/*!*/ /*!*/);
  --tw-sepia: var(--tw-empty,/*!*/ /*!*/);
  --tw-drop-shadow: var(--tw-empty,/*!*/ /*!*/);
  filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
}

.vue3-player-video .blur-sm{
  --tw-blur: blur(4px);
}

.vue3-player-video .blur{
  --tw-blur: blur(8px);
}

.vue3-player-video .backdrop-filter{
  --tw-backdrop-blur: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-brightness: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-contrast: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-grayscale: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-hue-rotate: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-invert: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-opacity: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-saturate: var(--tw-empty,/*!*/ /*!*/);
  --tw-backdrop-sepia: var(--tw-empty,/*!*/ /*!*/);
  -webkit-backdrop-filter: var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia);
          backdrop-filter: var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia);
}

.overlay {
  background: linear-gradient(0deg, rgba(0,0,0,0.41961), transparent)
}

.vertical-range::-webkit-slider-thumb {
  width: 6px;
  -webkit-appearance: none;
  appearance: none;
  height: 6px;
  background-color: white;
  cursor: ns-resize;
  box-shadow: -405px 0 0 400px rgba(255, 255, 255, .6);
  border-radius: 50%;
}

.backdrop-filter {
  -webkit-backdrop-filter: blur(15px) !important;
  backdrop-filter: blur(15px) !important;
}

.filter-white:hover {
  filter: brightness(2);
}

.gradient-variable {
  --tw-gradient-from: #fbbf24;
  --tw-gradient-to: #ec4899;
  --tw-gradient-stops: var(--tw-gradient-from), var(--tw-gradient-to, rgba(251, 191, 36, 0))
}

`;kn(xn);re.render=bn;var $n=(()=>{const e=re;return e.install=n=>{n.component("Vue3PlayerVideo",e)},e})(),Cn=$n;const Sn={key:0},Mn=$({props:{videos:{default:()=>[]},full:{type:Boolean,default:!1}},setup(e){const n=e;return(l,a)=>{const u=J,r=ee;return n.videos.length>0?(s(),c("div",Sn,[i(r,{"x-gap":4,"y-gap":4,cols:e.full?1:5},{default:p(()=>[i(u,{span:e.full?1:3},{default:p(()=>[(s(!0),c(h,null,k(n.videos,t=>(s(),y(m(Cn),{onClick:a[0]||(a[0]=x(()=>{},["stop"])),key:t.id,src:t.content,colors:["#18a058","#2aca75"],hoverable:!0,theme:"gradient"},null,8,["src"]))),128))]),_:1},8,["span"])]),_:1},8,["cols"])])):g("",!0)}}});const En={class:"images-wrap"},Nn=$({props:{imgs:{default:()=>[]}},setup(e){const n=e,l="https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/404.png",a="?x-oss-process=image/resize,m_fill,w_300,h_300,limit_0/auto-orient,1/format,png";return(u,r)=>{const t=we,d=J,v=ee,w=be;return s(),c("div",En,[[1].includes(n.imgs.length)?(s(),y(w,{key:0},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:2},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,o=>(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[0]||(r[0]=x(()=>{},["stop"])),class:"post-img x1","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024))),128))]),_:1})]),_:1})):g("",!0),[2,3].includes(n.imgs.length)?(s(),y(w,{key:1},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,o=>(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[1]||(r[1]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024))),128))]),_:1})]),_:1})):g("",!0),[4].includes(n.imgs.length)?(s(),y(w,{key:2},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:4},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,o=>(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[2]||(r[2]=x(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024))),128))]),_:1})]),_:1})):g("",!0),[5].includes(n.imgs.length)?(s(),y(w,{key:3},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,{key:o.id},[f<3?(s(),y(d,{key:0},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[3]||(r[3]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),128))]),_:1}),i(v,{"x-gap":4,"y-gap":4,cols:2,style:{"margin-top":"4px"}},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,{key:o.id},[f>=3?(s(),y(d,{key:0},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[4]||(r[4]=x(()=>{},["stop"])),class:"post-img x1","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),128))]),_:1})]),_:1})):g("",!0),[6].includes(n.imgs.length)?(s(),y(w,{key:4},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,{key:o.id},[f<3?(s(),y(d,{key:0},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[5]||(r[5]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),128))]),_:1}),i(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,{key:o.id},[f>=3?(s(),y(d,{key:0},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[6]||(r[6]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),128))]),_:1})]),_:1})):g("",!0),n.imgs.length===7?(s(),y(w,{key:5},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:4},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,null,[f<4?(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[7]||(r[7]=x(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),256))]),_:1}),i(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,null,[f>=4?(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[8]||(r[8]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),256))]),_:1})]),_:1})):g("",!0),n.imgs.length===8?(s(),y(w,{key:6},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:4},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,null,[f<4?(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[9]||(r[9]=x(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),256))]),_:1}),i(v,{"x-gap":4,"y-gap":4,cols:4,style:{"margin-top":"4px"}},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,null,[f>=4?(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[10]||(r[10]=x(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),256))]),_:1})]),_:1})):g("",!0),n.imgs.length===9?(s(),y(w,{key:7},{default:p(()=>[i(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,null,[f<3?(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[11]||(r[11]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),256))]),_:1}),i(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,null,[f>=3&&f<6?(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[12]||(r[12]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),256))]),_:1}),i(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(s(!0),c(h,null,k(n.imgs,(o,f)=>(s(),c(h,null,[f>=6?(s(),y(d,{key:o.id},{default:p(()=>[i(t,{onError:()=>o.content=m(l),onClick:r[13]||(r[13]=x(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:o.content+m(a),"preview-src":o.content},null,8,["onError","src","preview-src"])]),_:2},1024)):g("",!0)],64))),256))]),_:1})]),_:1})):g("",!0)])}}});const _n={class:"attachment-wrap"},Bn=$({props:{attachments:{default:()=>[]},price:{default:0}},setup(e){const n=e,l=I(!1),a=I(""),u=I(0),r=d=>{l.value=!0,u.value=d.id,a.value="\u8FD9\u662F\u4E00\u4E2A\u514D\u8D39\u9644\u4EF6\uFF0C\u60A8\u53EF\u4EE5\u76F4\u63A5\u4E0B\u8F7D\uFF1F",d.type===8&&(a.value=()=>R("div",{},[R("p",{},"\u8FD9\u662F\u4E00\u4E2A\u6536\u8D39\u9644\u4EF6\uFF0C\u4E0B\u8F7D\u5C06\u6536\u53D6"+(n.price/100).toFixed(2)+"\u5143")]),xe({id:u.value}).then(v=>{v.paid&&(a.value=()=>R("div",{},[R("p",{},"\u6B64\u6B21\u4E0B\u8F7D\u60A8\u5DF2\u652F\u4ED8\u6216\u65E0\u9700\u4ED8\u8D39\uFF0C\u8BF7\u786E\u8BA4\u4E0B\u8F7D")]))}).catch(v=>{l.value=!1}))},t=()=>{$e({id:u.value}).then(d=>{window.open(d.replace("http://","https://"),"_blank")}).catch(d=>{console.log(d)})};return(d,v)=>{const w=Y,o=Ce,f=Se;return s(),c("div",_n,[(s(!0),c(h,null,k(e.attachments,E=>(s(),c("div",{class:"attach-item",key:E.id},[i(o,{onClick:x(z=>r(E),["stop"]),type:"primary",size:"tiny",dashed:""},{icon:p(()=>[i(w,null,{default:p(()=>[i(m(Ue))]),_:1})]),default:p(()=>[ke(" "+N(E.type===8?"\u6536\u8D39":"\u514D\u8D39")+"\u9644\u4EF6 ",1)]),_:2},1032,["onClick"])]))),128)),i(f,{show:l.value,"onUpdate:show":v[0]||(v[0]=E=>l.value=E),"mask-closable":!1,preset:"dialog",title:"\u4E0B\u8F7D\u63D0\u793A",content:a.value,"positive-text":"\u786E\u8BA4\u4E0B\u8F7D","negative-text":"\u53D6\u6D88","icon-placement":"top",onPositiveClick:t},null,8,["show","content"])])}}});var An=H(Bn,[["__scopeId","data-v-ca4bb796"]]);const Pn=e=>{const n=[],l=[];var a=/(#|＃)([^#@])+?\s+?/g,u=/@([a-zA-Z0-9])+?\s+?/g;return e=e.replace(/<[^>]*?>/gi,"").replace(/(.*?)<\/[^>]*?>/gi,"").replace(a,r=>(n.push(r.substr(1).trim()),'<a class="hash-link" data-detail="tag:'+encodeURIComponent(r.substr(1).trim())+'">'+r.trim()+"</a> ")).replace(u,r=>(l.push(r.substr(1).trim()),'<a class="hash-link" data-detail="user:'+encodeURIComponent(r.substr(1).trim())+'">'+r.trim()+"</a> ")),{content:e,tags:n,users:l}};export{jn as B,Rn as C,Fn as H,An as _,Nn as a,Mn as b,In as c,Pn as p};
