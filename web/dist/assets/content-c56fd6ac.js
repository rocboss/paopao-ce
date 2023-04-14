import{bo as F,bp as pe,y as I,r as B,bq as ce,n as ve,d as C,q as fe,br as V,h as E,bs as me,u as ye,v as L,a2 as he,p as ge,t as Q,aU as we,b7 as W,bt as be,bu as ke,C as $e,D as xe,bv as Z,W as a,Y as c,Z as R,ab as g,ac as k,a4 as s,a5 as p,a3 as m,aa as G,a8 as $,af as J,al as ee,a6 as y,bw as q,bx as ne,a7 as w,by as te,bz as Se,bA as Ce,bB as _e,a9 as Re,bC as ze,bD as Ee,K as Be,aN as je}from"./index-cae59503.js";function Me(e){if(typeof e=="number")return{"":e.toString()};const n={};return e.split(/ +/).forEach(l=>{if(l==="")return;const[i,u]=l.split(":");u===void 0?n[""]=i:n[i]=u}),n}function O(e,n){var l;if(e==null)return;const i=Me(e);if(n===void 0)return i[""];if(typeof n=="string")return(l=i[n])!==null&&l!==void 0?l:i[""];if(Array.isArray(n)){for(let u=n.length-1;u>=0;--u){const o=n[u];if(o in i)return i[o]}return i[""]}else{let u,o=-1;return Object.keys(i).forEach(t=>{const d=Number(t);!Number.isNaN(d)&&n>=d&&d>=o&&(o=d,u=i[t])}),u}}function Ie(e){var n;const l=(n=e.dirs)===null||n===void 0?void 0:n.find(({dir:i})=>i===F);return!!(l&&l.value===!1)}const Ne={xs:0,s:640,m:1024,l:1280,xl:1536,"2xl":1920};function Te(e){return`(min-width: ${e}px)`}const U={};function Pe(e=Ne){if(!pe)return I(()=>[]);if(typeof window.matchMedia!="function")return I(()=>[]);const n=B({}),l=Object.keys(e),i=(u,o)=>{u.matches?n.value[o]=!0:n.value[o]=!1};return l.forEach(u=>{const o=e[u];let t,d;U[o]===void 0?(t=window.matchMedia(Te(o)),t.addEventListener?t.addEventListener("change",v=>{d.forEach(h=>{h(v,u)})}):t.addListener&&t.addListener(v=>{d.forEach(h=>{h(v,u)})}),d=new Set,U[o]={mql:t,cbs:d}):(t=U[o].mql,d=U[o].cbs),d.add(i),t.matches&&d.forEach(v=>{v(t,u)})}),ce(()=>{l.forEach(u=>{const{cbs:o}=U[e[u]];o.has(i)&&o.delete(i)})}),I(()=>{const{value:u}=n;return l.filter(o=>u[o])})}const K=1,re=ve("n-grid"),oe=1,Ae={span:{type:[Number,String],default:oe},offset:{type:[Number,String],default:0},suffix:Boolean,privateOffset:Number,privateSpan:Number,privateColStart:Number,privateShow:{type:Boolean,default:!0}},se=C({__GRID_ITEM__:!0,name:"GridItem",alias:["Gi"],props:Ae,setup(){const{isSsrRef:e,xGapRef:n,itemStyleRef:l,overflowRef:i,layoutShiftDisabledRef:u}=fe(re),o=me();return{overflow:i,itemStyle:l,layoutShiftDisabled:u,mergedXGap:I(()=>V(n.value||0)),deriveStyle:()=>{e.value;const{privateSpan:t=oe,privateShow:d=!0,privateColStart:v=void 0,privateOffset:h=0}=o.vnode.props,{value:r}=n,f=V(r||0);return{display:d?"":"none",gridColumn:`${v??`span ${t}`} / span ${t}`,marginLeft:h?`calc((100% - (${t} - 1) * ${f}) / ${t} * ${h} + ${f} * ${h})`:""}}}},render(){var e,n;if(this.layoutShiftDisabled){const{span:l,offset:i,mergedXGap:u}=this;return E("div",{style:{gridColumn:`span ${l} / span ${l}`,marginLeft:i?`calc((100% - (${l} - 1) * ${u}) / ${l} * ${i} + ${u} * ${i})`:""}},this.$slots)}return E("div",{style:[this.itemStyle,this.deriveStyle()]},(n=(e=this.$slots).default)===null||n===void 0?void 0:n.call(e,{overflow:this.overflow}))}}),Ve={xs:0,s:640,m:1024,l:1280,xl:1536,xxl:1920},ie=24,H="__ssr__",qe={layoutShiftDisabled:Boolean,responsive:{type:[String,Boolean],default:"self"},cols:{type:[Number,String],default:ie},itemResponsive:Boolean,collapsed:Boolean,collapsedRows:{type:Number,default:1},itemStyle:[Object,String],xGap:{type:[Number,String],default:0},yGap:{type:[Number,String],default:0}},ae=C({name:"Grid",inheritAttrs:!1,props:qe,setup(e){const{mergedClsPrefixRef:n,mergedBreakpointsRef:l}=ye(e),i=/^\d+$/,u=B(void 0),o=Pe((l==null?void 0:l.value)||Ve),t=L(()=>!!(e.itemResponsive||!i.test(e.cols.toString())||!i.test(e.xGap.toString())||!i.test(e.yGap.toString()))),d=I(()=>{if(t.value)return e.responsive==="self"?u.value:o.value}),v=L(()=>{var x;return(x=Number(O(e.cols.toString(),d.value)))!==null&&x!==void 0?x:ie}),h=L(()=>O(e.xGap.toString(),d.value)),r=L(()=>O(e.yGap.toString(),d.value)),f=x=>{u.value=x.contentRect.width},S=x=>{ke(f,x)},N=B(!1),T=I(()=>{if(e.responsive==="self")return S}),_=B(!1),j=B();return he(()=>{const{value:x}=j;x&&x.hasAttribute(H)&&(x.removeAttribute(H),_.value=!0)}),ge(re,{layoutShiftDisabledRef:Q(e,"layoutShiftDisabled"),isSsrRef:_,itemStyleRef:Q(e,"itemStyle"),xGapRef:h,overflowRef:N}),{isSsr:!we,contentEl:j,mergedClsPrefix:n,style:I(()=>e.layoutShiftDisabled?{width:"100%",display:"grid",gridTemplateColumns:`repeat(${e.cols}, minmax(0, 1fr))`,columnGap:V(e.xGap),rowGap:V(e.yGap)}:{width:"100%",display:"grid",gridTemplateColumns:`repeat(${v.value}, minmax(0, 1fr))`,columnGap:V(h.value),rowGap:V(r.value)}),isResponsive:t,responsiveQuery:d,responsiveCols:v,handleResize:T,overflow:N}},render(){if(this.layoutShiftDisabled)return E("div",W({ref:"contentEl",class:`${this.mergedClsPrefix}-grid`,style:this.style},this.$attrs),this.$slots);const e=()=>{var n,l,i,u,o,t,d;this.overflow=!1;const v=$e(xe(this)),h=[],{collapsed:r,collapsedRows:f,responsiveCols:S,responsiveQuery:N}=this;v.forEach(b=>{var D,M,z,P;if(((D=b==null?void 0:b.type)===null||D===void 0?void 0:D.__GRID_ITEM__)!==!0)return;if(Ie(b)){const A=Z(b);A.props?A.props.privateShow=!1:A.props={privateShow:!1},h.push({child:A,rawChildSpan:0});return}b.dirs=((M=b.dirs)===null||M===void 0?void 0:M.filter(({dir:A})=>A!==F))||null;const X=Z(b),Y=Number((P=O((z=X.props)===null||z===void 0?void 0:z.span,N))!==null&&P!==void 0?P:K);Y!==0&&h.push({child:X,rawChildSpan:Y})});let T=0;const _=(n=h[h.length-1])===null||n===void 0?void 0:n.child;if(_!=null&&_.props){const b=(l=_.props)===null||l===void 0?void 0:l.suffix;b!==void 0&&b!==!1&&(T=(u=(i=_.props)===null||i===void 0?void 0:i.span)!==null&&u!==void 0?u:K,_.props.privateSpan=T,_.props.privateColStart=S+1-T,_.props.privateShow=(o=_.props.privateShow)!==null&&o!==void 0?o:!0)}let j=0,x=!1;for(const{child:b,rawChildSpan:D}of h){if(x&&(this.overflow=!0),!x){const M=Number((d=O((t=b.props)===null||t===void 0?void 0:t.offset,N))!==null&&d!==void 0?d:0),z=Math.min(D+M,S);if(b.props?(b.props.privateSpan=z,b.props.privateOffset=M):b.props={privateSpan:z,privateOffset:M},r){const P=j%S;z+P>S&&(j+=S-P),z+j+T>f*S?x=!0:j+=z}}x&&(b.props?b.props.privateShow!==!0&&(b.props.privateShow=!1):b.props={privateShow:!1})}return E("div",W({ref:"contentEl",class:`${this.mergedClsPrefix}-grid`,style:this.style,[H]:this.isSsr||void 0},this.$attrs),h.map(({child:b})=>b))};return this.isResponsive&&this.responsive==="self"?E(be,{onResize:this.handleResize},{default:e}):e()}}),Fe={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Ge=R("path",{d:"M352 48H160a48 48 0 0 0-48 48v368l144-128l144 128V96a48 48 0 0 0-48-48z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),De=[Ge],An=C({name:"BookmarkOutline",render:function(n,l){return a(),c("svg",Fe,De)}}),Oe={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Ue=R("path",{d:"M408 64H104a56.16 56.16 0 0 0-56 56v192a56.16 56.16 0 0 0 56 56h40v80l93.72-78.14a8 8 0 0 1 5.13-1.86H408a56.16 56.16 0 0 0 56-56V120a56.16 56.16 0 0 0-56-56z",fill:"none",stroke:"currentColor","stroke-linejoin":"round","stroke-width":"32"},null,-1),Le=[Ue],Vn=C({name:"ChatboxOutline",render:function(n,l){return a(),c("svg",Oe,Le)}}),He={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Xe=R("path",{d:"M320 336h76c55 0 100-21.21 100-75.6s-53-73.47-96-75.6C391.11 99.74 329 48 256 48c-69 0-113.44 45.79-128 91.2c-60 5.7-112 35.88-112 98.4S70 336 136 336h56",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),Ye=R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M192 400.1l64 63.9l64-63.9"},null,-1),Qe=R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M256 224v224.03"},null,-1),We=[Xe,Ye,Qe],Ze=C({name:"CloudDownloadOutline",render:function(n,l){return a(),c("svg",He,We)}}),Ke={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Je=R("path",{d:"M352.92 80C288 80 256 144 256 144s-32-64-96.92-64c-52.76 0-94.54 44.14-95.08 96.81c-1.1 109.33 86.73 187.08 183 252.42a16 16 0 0 0 18 0c96.26-65.34 184.09-143.09 183-252.42c-.54-52.67-42.32-96.81-95.08-96.81z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),en=[Je],qn=C({name:"HeartOutline",render:function(n,l){return a(),c("svg",Ke,en)}}),nn={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},tn=R("path",{d:"M208 352h-64a96 96 0 0 1 0-192h64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),rn=R("path",{d:"M304 160h64a96 96 0 0 1 0 192h-64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),on=R("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36",d:"M163.29 256h187.42"},null,-1),sn=[tn,rn,on],an=C({name:"LinkOutline",render:function(n,l){return a(),c("svg",nn,sn)}}),ln={class:"link-wrap"},un=["href"],dn={class:"link-txt"},pn=C({__name:"post-link",props:{links:{default:()=>[]}},setup(e){const n=e;return(l,i)=>{const u=J;return a(),c("div",ln,[(a(!0),c(g,null,k(n.links,o=>(a(),c("div",{class:"link-item",key:o.id},[s(u,{class:"hash-link"},{default:p(()=>[s(m(an))]),_:1}),R("a",{href:o.content,class:"hash-link",target:"_blank",onClick:i[0]||(i[0]=$(()=>{},["stop"]))},[R("span",dn,G(o.content),1)],8,un)]))),128))])}}});const Fn=ee(pn,[["__scopeId","data-v-6c4d1eb6"]]);var le=C({name:"BasicTheme",props:{uuid:{type:String,required:!0},src:{type:String,required:!0},autoplay:{type:Boolean,required:!0},loop:{type:Boolean,required:!0},controls:{type:Boolean,required:!0},hoverable:{type:Boolean,required:!0},mask:{type:Boolean,required:!0},colors:{type:[String,Array],required:!0},time:{type:Object,required:!0},playing:{type:Boolean,default:!1},duration:{type:[String,Number],required:!0}},data(){return{hovered:!1,volume:!1,amount:1}},computed:{colorFrom(){var e;return typeof this.colors=="string"?this.colors?this.colors:"#fbbf24":(e=this.colors)!==null&&e!==void 0&&e[0]?this.colors[0]:"#fbbf24"},colorTo(){var e;return typeof this.colors=="string"?this.colors?this.colors:"#fbbf24":(e=this.colors)!==null&&e!==void 0&&e[1]?this.colors[1]:"#ec4899"}},mounted(){this.$emit("setPlayer",this.$refs[this.uuid])},methods:{setVolume(){this.$refs[this.uuid].volume=this.amount},stopVolume(){return this.amount>0?this.amount=0:this.amount=1}}});const cn={class:"relative"},vn={class:"flex items-center justify-start w-full"},fn={class:"font-sans text-white text-xs w-24"},mn={class:"mr-3 ml-2"},yn={class:"relative"},hn={class:"px-3 py-2 rounded-lg flex items-center transform translate-x-2",style:{"background-color":"rgba(0, 0, 0, .8)"}},gn=s("img",{src:"https://en-zo.dev/vue-videoplayer/play.svg",alt:"Icon play video",class:"transform translate-x-0.5 w-12"},null,-1);function wn(e,n,l,i,u,o){return a(),y("div",{class:"shadow-xl rounded-xl overflow-hidden relative",onMouseenter:n[15]||(n[15]=t=>e.hovered=!0),onMouseleave:n[16]||(n[16]=t=>e.hovered=!1),onKeydown:n[17]||(n[17]=te(t=>e.$emit("play"),["left"]))},[s("div",cn,[s("video",{ref:e.uuid,class:"w-full",loop:e.loop,autoplay:e.autoplay,muted:e.autoplay,onTimeupdate:n[1]||(n[1]=t=>e.$emit("timeupdate",t.target)),onPause:n[2]||(n[2]=t=>e.$emit("isPlaying",!1)),onPlay:n[3]||(n[3]=t=>e.$emit("isPlaying",!0)),onClick:n[4]||(n[4]=t=>e.$emit("play"))},[s("source",{src:e.src,type:"video/mp4"},null,8,["src"])],40,["loop","autoplay","muted"]),e.controls?(a(),y("div",{key:0,class:[{"opacity-0 translate-y-full":!e.hoverable&&e.hovered,"opacity-0 translate-y-full":e.hoverable&&!e.hovered},"transition duration-300 transform absolute w-full bottom-0 left-0 flex items-center justify-between overlay px-5  pt-3 pb-5"]},[s("div",vn,[s("p",fn,G(e.time.display)+"/"+G(e.duration),1),s("div",mn,[q(s("img",{src:"https://en-zo.dev/vue-videoplayer/pause.svg",alt:"Icon pause video",class:"w-5 cursor-pointer",onClick:n[5]||(n[5]=t=>e.$emit("play"))},null,512),[[F,e.playing]]),q(s("img",{src:"https://en-zo.dev/vue-videoplayer/play.svg",alt:"Icon play video",class:"w-5 cursor-pointer",onClick:n[6]||(n[6]=t=>e.$emit("play"))},null,512),[[F,!e.playing]])]),s("div",{class:"w-full h-1 bg-white bg-opacity-60 rounded-sm cursor-pointer",onClick:n[7]||(n[7]=t=>e.$emit("position",t))},[s("div",{class:"relative h-full pointer-events-none",style:`width: ${e.time.progress}%; transition: width .2s ease-in-out;`},[s("div",{class:"w-full rounded-sm h-full gradient-variable bg-gradient-to-r pointer-events-none absolute top-0 left-0",style:`--tw-gradient-from: ${e.colorFrom};--tw-gradient-to: ${e.colorTo};transition: width .2s ease-in-out`},null,4),s("div",{class:"w-full rounded-sm filter blur-sm h-full gradient-variable bg-gradient-to-r pointer-events-none absolute top-0 left-0",style:`--tw-gradient-from: ${e.colorFrom};--tw-gradient-to: ${e.colorTo};transition: width .2s ease-in-out`},null,4)],4)])]),s("div",{class:"ml-5 flex items-center justify-end",onMouseleave:n[13]||(n[13]=t=>e.volume=!1)},[s("div",yn,[s("div",{class:`w-128 origin-left translate-x-2 -rotate-90 w-128 transition duration-200 absolute transform top-0 py-2 ${e.volume?"-translate-y-4":"opacity-0 -translate-y-1 pointer-events-none"}`},[s("div",hn,[q(s("input",{"onUpdate:modelValue":n[8]||(n[8]=t=>e.amount=t),type:"range",step:"0.05",min:"0",max:"1",class:"rounded-lg overflow-hidden appearance-none bg-white bg-opacity-30 h-1 w-128 vertical-range",onInput:n[9]||(n[9]=(...t)=>e.setVolume&&e.setVolume(...t))},null,544),[[ne,e.amount]])])],2),s("img",{src:`https://en-zo.dev/vue-videoplayer/volume-${Math.ceil(e.amount*2)}.svg`,alt:"High volume video",class:"w-5 cursor-pointer relative",style:{"z-index":"2"},onClick:n[10]||(n[10]=(...t)=>e.stopVolume&&e.stopVolume(...t)),onMouseenter:n[11]||(n[11]=t=>e.volume=!0)},null,40,["src"])]),s("img",{src:"https://en-zo.dev/vue-videoplayer/maximize.svg",alt:"Fullscreen",class:"w-3 ml-4 cursor-pointer",onClick:n[12]||(n[12]=t=>e.$emit("fullScreen"))})],32)],2)):w("",!0),!e.autoplay&&e.mask&&e.time.current===0?(a(),y("div",{key:1,class:`transition transform duration-300 absolute top-0 left-0 w-full h-full bg-black bg-opacity-50 backdrop-filter z-10 flex items-center justify-center ${e.playing?"opacity-0 pointer-events-none":""}`},[s("div",{class:"w-20 h-20 rounded-full bg-white bg-opacity-20 transition duration-200 hover:bg-opacity-40 flex items-center justify-center cursor-pointer",onClick:n[14]||(n[14]=t=>e.$emit("play"))},[gn])],2)):w("",!0)])],32)}le.render=wn;var ue=C({name:"BasicTheme",props:{uuid:{type:String,required:!0},src:{type:String,required:!0},autoplay:{type:Boolean,required:!0},loop:{type:Boolean,required:!0},controls:{type:Boolean,required:!0},hoverable:{type:Boolean,required:!0},mask:{type:Boolean,required:!0},colors:{type:[String,Array],required:!0},time:{type:Object,required:!0},playing:{type:Boolean,default:!1},duration:{type:[String,Number],required:!0}},data(){return{hovered:!1,volume:!1,amount:1}},computed:{color(){var e;return typeof this.colors=="string"?this.colors?this.colors:"#8B5CF6":(e=this.colors)!==null&&e!==void 0&&e[0]?this.colors[0]:"#8B5CF6"}},mounted(){this.$emit("setPlayer",this.$refs[this.uuid])},methods:{setVolume(){this.$refs[this.uuid].volume=this.amount},stopVolume(){return this.amount>0?this.amount=0:this.amount=1}}});const bn={class:"relative"},kn={class:"mr-5"},$n={class:"relative mr-6"},xn={class:"px-3 py-3 rounded-xl flex items-center transform translate-x-9 bg-black bg-opacity-30"},Sn=s("img",{src:"https://en-zo.dev/vue-videoplayer/play.svg",alt:"Icon play video",class:"transform translate-x-0.5 w-12"},null,-1);function Cn(e,n,l,i,u,o){return a(),y("div",{class:"shadow-xl rounded-3xl overflow-hidden relative",onMouseenter:n[14]||(n[14]=t=>e.hovered=!0),onMouseleave:n[15]||(n[15]=t=>e.hovered=!1),onKeydown:n[16]||(n[16]=te(t=>e.$emit("play"),["left"]))},[s("div",bn,[s("video",{ref:e.uuid,class:"w-full",loop:e.loop,autoplay:e.autoplay,muted:e.autoplay,onTimeupdate:n[1]||(n[1]=t=>e.$emit("timeupdate",t.target)),onPause:n[2]||(n[2]=t=>e.$emit("isPlaying",!1)),onPlay:n[3]||(n[3]=t=>e.$emit("isPlaying",!0)),onClick:n[4]||(n[4]=t=>e.$emit("play"))},[s("source",{src:e.src,type:"video/mp4"},null,8,["src"])],40,["loop","autoplay","muted"]),e.controls?(a(),y("div",{key:0,class:[{"opacity-0 translate-y-full":!e.hoverable&&e.hovered,"opacity-0 translate-y-full":e.hoverable&&!e.hovered},"absolute px-5 pb-5 bottom-0 left-0 w-full transition duration-300 transform"]},[s("div",{class:"w-full bg-black bg-opacity-30 px-5 py-4 rounded-xl flex items-center justify-between",onMouseleave:n[12]||(n[12]=t=>e.volume=!1)},[s("div",{class:"font-sans py-1 px-2 text-white rounded-md text-xs mr-5 whitespace-nowrap font-medium w-32 text-center",style:`font-size: 11px; background-color: ${e.color}`},G(e.time.display)+" / "+G(e.duration),5),s("div",kn,[q(s("img",{src:"https://en-zo.dev/vue-videoplayer/basic/pause.svg",alt:"Icon pause video",class:"w-4 cursor-pointer filter-white transition duration-300",onClick:n[5]||(n[5]=t=>e.$emit("play"))},null,512),[[F,e.playing]]),q(s("img",{src:"https://en-zo.dev/vue-videoplayer/basic/play.svg",alt:"Icon play video",class:"w-4 cursor-pointer filter-white transition duration-300",onClick:n[6]||(n[6]=t=>e.$emit("play"))},null,512),[[F,!e.playing]])]),s("div",{class:"w-full h-1 bg-white bg-opacity-40 rounded-sm cursor-pointer mr-6",onClick:n[7]||(n[7]=t=>e.$emit("position",t))},[s("div",{class:"w-full rounded-sm h-full bg-white pointer-events-none",style:`width: ${e.time.progress}%; transition: width .2s ease-in-out;`},null,4)]),s("div",$n,[s("div",{class:`w-128 origin-left translate-x-2 -rotate-90 w-128 transition duration-200 absolute transform top-0 py-2 ${e.volume?"-translate-y-4":"opacity-0 -translate-y-1 pointer-events-none"}`},[s("div",xn,[q(s("input",{"onUpdate:modelValue":n[8]||(n[8]=t=>e.amount=t),type:"range",step:"0.05",min:"0",max:"1",class:"rounded-lg overflow-hidden appearance-none bg-white bg-opacity-30 h-1.5 w-128 vertical-range"},null,512),[[ne,e.amount]])])],2),s("img",{src:`https://en-zo.dev/vue-videoplayer/basic/volume_${Math.ceil(e.amount*2)}.svg`,alt:"High volume video",class:"w-5 cursor-pointer filter-white transition duration-300 relative",style:{"z-index":"2"},onClick:n[9]||(n[9]=(...t)=>e.stopVolume&&e.stopVolume(...t)),onMouseenter:n[10]||(n[10]=t=>e.volume=!0)},null,40,["src"])]),s("img",{src:"https://en-zo.dev/vue-videoplayer/basic/fullscreen.svg",alt:"Fullscreen",class:"w-4 cursor-pointer filter-white transition duration-300",onClick:n[11]||(n[11]=t=>e.$emit("fullScreen"))})],32)],2)):w("",!0),!e.autoplay&&e.mask&&e.time.current===0?(a(),y("div",{key:1,class:`transition transform duration-300 absolute top-0 left-0 w-full h-full bg-black bg-opacity-50 backdrop-filter z-10 flex items-center justify-center ${e.playing?"opacity-0 pointer-events-none":""}`},[s("div",{class:"w-20 h-20 rounded-full bg-white bg-opacity-20 transition duration-200 hover:bg-opacity-40 flex items-center justify-center cursor-pointer",onClick:n[13]||(n[13]=t=>e.$emit("play"))},[Sn])],2)):w("",!0)])],32)}ue.render=Cn;var de=C({name:"Vue3PlayerVideo",components:{basic:ue,gradient:le},props:{src:{type:String,required:!0},autoplay:{type:Boolean,default:!1},loop:{type:Boolean,default:!1},controls:{type:Boolean,default:!0},mask:{type:Boolean,default:!0},colors:{type:[String,Array],default(){return["#8B5CF6","#ec4899"]}},hoverable:{type:Boolean,default:!1},theme:{type:String,default:"basic"}},data(){return{uuid:Math.random().toString(36).substr(2,18),player:null,duration:0,playing:!1,time:{progress:0,display:0,current:0}}},watch:{"time.current"(e){this.time.display=this.format(Number(e)),this.time.progress=e*100/this.player.duration}},methods:{isPlaying(e){this.playing=e},play(){return this.playing?this.player.pause():this.player.play()},setPlayer(e){this.player=e,this.player.addEventListener("loadeddata",()=>{this.player.readyState>=3&&(this.duration=this.format(Number(this.player.duration)),this.time.display=this.format(0))})},stop(){this.player.pause(),this.player.currentTime=0},fullScreen(){this.player.webkitEnterFullscreen()},position(e){this.player.pause();const n=e.target.getBoundingClientRect(),i=(e.clientX-n.left)*100/e.target.offsetWidth;this.player.currentTime=i*this.player.duration/100,this.player.play()},format(e){const n=Math.floor(e/3600),l=Math.floor(e%3600/60),i=Math.round(e%60);return[n,l>9?l:n?"0"+l:l||"00",i>9?i:"0"+i].filter(Boolean).join(":")}}});const _n={class:"vue3-player-video"};function Rn(e,n,l,i,u,o){return a(),y("div",_n,[(a(),y(Se(e.theme),{uuid:e.uuid,src:e.src,autoplay:e.autoplay,loop:e.loop,controls:e.controls,mask:e.mask,colors:e.colors,time:e.time,playing:e.playing,duration:e.duration,hoverable:e.hoverable,onPlay:e.play,onStop:e.stop,onTimeupdate:n[1]||(n[1]=({currentTime:t})=>e.time.current=t),onPosition:e.position,onFullScreen:e.fullScreen,onSetPlayer:e.setPlayer,onIsPlaying:e.isPlaying},null,8,["uuid","src","autoplay","loop","controls","mask","colors","time","playing","duration","hoverable","onPlay","onStop","onPosition","onFullScreen","onSetPlayer","onIsPlaying"]))])}function zn(e,n){n===void 0&&(n={});var l=n.insertAt;if(!(!e||typeof document>"u")){var i=document.head||document.getElementsByTagName("head")[0],u=document.createElement("style");u.type="text/css",l==="top"&&i.firstChild?i.insertBefore(u,i.firstChild):i.appendChild(u),u.styleSheet?u.styleSheet.cssText=e:u.appendChild(document.createTextNode(e))}}var En=`/*! tailwindcss v2.1.2 | MIT License | https://tailwindcss.com */

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

`;zn(En);de.render=Rn;var Bn=(()=>{const e=de;return e.install=n=>{n.component("Vue3PlayerVideo",e)},e})();const jn=Bn,Mn={key:0},Gn=C({__name:"post-video",props:{videos:{default:()=>[]},full:{type:Boolean,default:!1}},setup(e){const n=e;return(l,i)=>{const u=se,o=ae;return n.videos.length>0?(a(),c("div",Mn,[s(o,{"x-gap":4,"y-gap":4,cols:e.full?1:5},{default:p(()=>[s(u,{span:e.full?1:3},{default:p(()=>[(a(!0),c(g,null,k(n.videos,t=>(a(),y(m(jn),{onClick:i[0]||(i[0]=$(()=>{},["stop"])),key:t.id,src:t.content,colors:["#18a058","#2aca75"],hoverable:!0,theme:"gradient"},null,8,["src"]))),128))]),_:1},8,["span"])]),_:1},8,["cols"])])):w("",!0)}}}),In={class:"images-wrap"},Dn=C({__name:"post-image",props:{imgs:{default:()=>[]}},setup(e){const n=e,l="https://paopao-assets.oss-cn-shanghai.aliyuncs.com/public/404.png",i="?x-oss-process=image/resize,m_fill,w_300,h_300,limit_0/auto-orient,1/format,png";return(u,o)=>{const t=Ce,d=se,v=ae,h=_e;return a(),c("div",In,[[1].includes(n.imgs.length)?(a(),y(h,{key:0},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:2},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,r=>(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[0]||(o[0]=$(()=>{},["stop"])),class:"post-img x1","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024))),128))]),_:1})]),_:1})):w("",!0),[2,3].includes(n.imgs.length)?(a(),y(h,{key:1},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,r=>(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[1]||(o[1]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024))),128))]),_:1})]),_:1})):w("",!0),[4].includes(n.imgs.length)?(a(),y(h,{key:2},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:4},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,r=>(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[2]||(o[2]=$(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024))),128))]),_:1})]),_:1})):w("",!0),[5].includes(n.imgs.length)?(a(),y(h,{key:3},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,{key:r.id},[f<3?(a(),y(d,{key:0},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[3]||(o[3]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),128))]),_:1}),s(v,{"x-gap":4,"y-gap":4,cols:2,style:{"margin-top":"4px"}},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,{key:r.id},[f>=3?(a(),y(d,{key:0},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[4]||(o[4]=$(()=>{},["stop"])),class:"post-img x1","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),128))]),_:1})]),_:1})):w("",!0),[6].includes(n.imgs.length)?(a(),y(h,{key:4},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,{key:r.id},[f<3?(a(),y(d,{key:0},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[5]||(o[5]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),128))]),_:1}),s(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,{key:r.id},[f>=3?(a(),y(d,{key:0},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[6]||(o[6]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),128))]),_:1})]),_:1})):w("",!0),n.imgs.length===7?(a(),y(h,{key:5},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:4},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,null,[f<4?(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[7]||(o[7]=$(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),256))]),_:1}),s(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,null,[f>=4?(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[8]||(o[8]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),256))]),_:1})]),_:1})):w("",!0),n.imgs.length===8?(a(),y(h,{key:6},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:4},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,null,[f<4?(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[9]||(o[9]=$(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),256))]),_:1}),s(v,{"x-gap":4,"y-gap":4,cols:4,style:{"margin-top":"4px"}},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,null,[f>=4?(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[10]||(o[10]=$(()=>{},["stop"])),class:"post-img x3","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),256))]),_:1})]),_:1})):w("",!0),n.imgs.length===9?(a(),y(h,{key:7},{default:p(()=>[s(v,{"x-gap":4,"y-gap":4,cols:3},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,null,[f<3?(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[11]||(o[11]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),256))]),_:1}),s(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,null,[f>=3&&f<6?(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[12]||(o[12]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),256))]),_:1}),s(v,{"x-gap":4,"y-gap":4,cols:3,style:{"margin-top":"4px"}},{default:p(()=>[(a(!0),c(g,null,k(n.imgs,(r,f)=>(a(),c(g,null,[f>=6?(a(),y(d,{key:r.id},{default:p(()=>[s(t,{onError:()=>r.content=m(l),onClick:o[13]||(o[13]=$(()=>{},["stop"])),class:"post-img x2","object-fit":"cover",src:r.content+m(i),"preview-src":r.content},null,8,["onError","src","preview-src"])]),_:2},1024)):w("",!0)],64))),256))]),_:1})]),_:1})):w("",!0)])}}});const Nn={class:"attachment-wrap"},Tn=C({__name:"post-attachment",props:{attachments:{default:()=>[]},price:{default:0}},setup(e){const n=e,l=B(!1),i=B(""),u=B(0),o=d=>{l.value=!0,u.value=d.id,i.value="这是一个免费附件，您可以直接下载？",d.type===8&&(i.value=()=>E("div",{},[E("p",{},"这是一个收费附件，下载将收取"+(n.price/100).toFixed(2)+"元")]),ze({id:u.value}).then(v=>{v.paid&&(i.value=()=>E("div",{},[E("p",{},"此次下载您已支付或无需付费，请确认下载")]))}).catch(v=>{l.value=!1}))},t=()=>{Ee({id:u.value}).then(d=>{window.open(d.signed_url.replace("http://","https://"),"_blank")}).catch(d=>{console.log(d)})};return(d,v)=>{const h=J,r=Be,f=je;return a(),c("div",Nn,[(a(!0),c(g,null,k(e.attachments,S=>(a(),c("div",{class:"attach-item",key:S.id},[s(r,{onClick:$(N=>o(S),["stop"]),type:"primary",size:"tiny",dashed:""},{icon:p(()=>[s(h,null,{default:p(()=>[s(m(Ze))]),_:1})]),default:p(()=>[Re(" "+G(S.type===8?"收费":"免费")+"附件 ",1)]),_:2},1032,["onClick"])]))),128)),s(f,{show:l.value,"onUpdate:show":v[0]||(v[0]=S=>l.value=S),"mask-closable":!1,preset:"dialog",title:"下载提示",content:i.value,"positive-text":"确认下载","negative-text":"取消","icon-placement":"top",onPositiveClick:t},null,8,["show","content"])])}}});const On=ee(Tn,[["__scopeId","data-v-22563084"]]),Un=e=>{const n=[],l=[];var i=/(#|＃)([^#@\s])+?\s+?/g,u=/@([a-zA-Z0-9])+?\s+?/g;return e=e.replace(/<[^>]*?>/gi,"").replace(/(.*?)<\/[^>]*?>/gi,"").replace(i,o=>(n.push(o.substr(1).trim()),'<a class="hash-link" data-detail="tag:'+encodeURIComponent(o.substr(1).trim())+'">'+o.trim()+"</a> ")).replace(u,o=>(l.push(o.substr(1).trim()),'<a class="hash-link" data-detail="user:'+encodeURIComponent(o.substr(1).trim())+'">'+o.trim()+"</a> ")),{content:e,tags:n,users:l}};export{An as B,Vn as C,qn as H,Dn as _,On as a,Gn as b,Fn as c,Un as p};
