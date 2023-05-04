import{ba as je,d as O,aE as De,aF as Ae,a2 as se,bb as Ke,bc as We,y as T,r as $,v as Y,aw as re,av as K,h as d,bd as ye,be as de,aH as ne,bf as Ge,bg as ce,as as xe,c as L,f as V,b as j,u as we,x as W,bh as qe,J as Ue,q as Z,bi as Xe,z as D,A as Se,N as ze,bj as ae,aJ as Re,b3 as Ye,e as A,a as Ze,aX as Je,t as H,bk as Qe,bl as et,S as ue,V as tt,bm as oe,p as fe,bn as nt,B as ot,az as it,bo as lt,L as rt,aS as ve,bp as st,k as at,ab as dt}from"./index.f912861b.js";import{l as ct}from"./List.28de093a.js";function ie(e){const t=e.filter(o=>o!==void 0);if(t.length!==0)return t.length===1?t[0]:o=>{e.forEach(s=>{s&&s(o)})}}let he=!1;function ut(){if(!!je&&!!window.CSS&&!he&&(he=!0,"registerProperty"in(window==null?void 0:window.CSS)))try{CSS.registerProperty({name:"--n-color-start",syntax:"<color>",inherits:!1,initialValue:"#0000"}),CSS.registerProperty({name:"--n-color-end",syntax:"<color>",inherits:!1,initialValue:"#0000"})}catch{}}function me(e){return e&-e}class ft{constructor(t,o){this.l=t,this.min=o;const s=new Array(t+1);for(let r=0;r<t+1;++r)s[r]=0;this.ft=s}add(t,o){if(o===0)return;const{l:s,ft:r}=this;for(t+=1;t<=s;)r[t]+=o,t+=me(t)}get(t){return this.sum(t+1)-this.sum(t)}sum(t){if(t===void 0&&(t=this.l),t<=0)return 0;const{ft:o,min:s,l:r}=this;if(t>r)throw new Error("[FinweckTree.sum]: `i` is larger than length.");let a=t*s;for(;t>0;)a+=o[t],t-=me(t);return a}getBound(t){let o=0,s=this.l;for(;s>o;){const r=Math.floor((o+s)/2),a=this.sum(r);if(a>t){s=r;continue}else if(a<t){if(o===r)return this.sum(o+1)<=t?o+1:r;o=r}else return r}return o}}let X;function vt(){return X===void 0&&("matchMedia"in window?X=window.matchMedia("(pointer:coarse)").matches:X=!1),X}let le;function ge(){return le===void 0&&(le="chrome"in window?window.devicePixelRatio:1),le}const ht=ne(".v-vl",{maxHeight:"inherit",height:"100%",overflow:"auto",minWidth:"1px"},[ne("&:not(.v-vl--show-scrollbar)",{scrollbarWidth:"none"},[ne("&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb",{width:0,height:0,display:"none"})])]);var mt=O({name:"VirtualList",inheritAttrs:!1,props:{showScrollbar:{type:Boolean,default:!0},items:{type:Array,default:()=>[]},itemSize:{type:Number,required:!0},itemResizable:Boolean,itemsStyle:[String,Object],visibleItemsTag:{type:[String,Object],default:"div"},visibleItemsProps:Object,ignoreItemResize:Boolean,onScroll:Function,onWheel:Function,onResize:Function,defaultScrollKey:[Number,String],defaultScrollIndex:Number,keyField:{type:String,default:"key"},paddingTop:{type:[Number,String],default:0},paddingBottom:{type:[Number,String],default:0}},setup(e){const t=De();ht.mount({id:"vueuc/virtual-list",head:!0,anchorMetaName:Ae,ssr:t}),se(()=>{const{defaultScrollIndex:i,defaultScrollKey:c}=e;i!=null?h({index:i}):c!=null&&h({key:c})});let o=!1,s=!1;Ke(()=>{if(o=!1,!s){s=!0;return}h({top:m.value,left:f})}),We(()=>{o=!0,s||(s=!0)});const r=T(()=>{const i=new Map,{keyField:c}=e;return e.items.forEach((g,k)=>{i.set(g[c],k)}),i}),a=$(null),u=$(void 0),v=new Map,R=T(()=>{const{items:i,itemSize:c,keyField:g}=e,k=new ft(i.length,c);return i.forEach((C,P)=>{const p=C[g],_=v.get(p);_!==void 0&&k.add(P,_)}),k}),b=$(0);let f=0;const m=$(0),z=Y(()=>Math.max(R.value.getBound(m.value-re(e.paddingTop))-1,0)),F=T(()=>{const{value:i}=u;if(i===void 0)return[];const{items:c,itemSize:g}=e,k=z.value,C=Math.min(k+Math.ceil(i/g+1),c.length-1),P=[];for(let p=k;p<=C;++p)P.push(c[p]);return P}),h=(i,c)=>{if(typeof i=="number"){x(i,c,"auto");return}const{left:g,top:k,index:C,key:P,position:p,behavior:_,debounce:n=!0}=i;if(g!==void 0||k!==void 0)x(g,k,_);else if(C!==void 0)S(C,_,n);else if(P!==void 0){const l=r.value.get(P);l!==void 0&&S(l,_,n)}else p==="bottom"?x(0,Number.MAX_SAFE_INTEGER,_):p==="top"&&x(0,0,_)};let y,M=null;function S(i,c,g){const{value:k}=R,C=k.sum(i)+re(e.paddingTop);if(!g)a.value.scrollTo({left:0,top:C,behavior:c});else{y=i,M!==null&&window.clearTimeout(M),M=window.setTimeout(()=>{y=void 0,M=null},16);const{scrollTop:P,offsetHeight:p}=a.value;if(C>P){const _=k.get(i);C+_<=P+p||a.value.scrollTo({left:0,top:C+_-p,behavior:c})}else a.value.scrollTo({left:0,top:C,behavior:c})}}function x(i,c,g){a.value.scrollTo({left:i,top:c,behavior:g})}function I(i,c){var g,k,C;if(o||e.ignoreItemResize||U(c.target))return;const{value:P}=R,p=r.value.get(i),_=P.get(p),n=(C=(k=(g=c.borderBoxSize)===null||g===void 0?void 0:g[0])===null||k===void 0?void 0:k.blockSize)!==null&&C!==void 0?C:c.contentRect.height;if(n===_)return;n-e.itemSize===0?v.delete(i):v.set(i,n-e.itemSize);const w=n-_;if(w===0)return;P.add(p,w);const N=a.value;if(N!=null){if(y===void 0){const q=P.sum(p);N.scrollTop>q&&N.scrollBy(0,w)}else if(p<y)N.scrollBy(0,w);else if(p===y){const q=P.sum(p);n+q>N.scrollTop+N.offsetHeight&&N.scrollBy(0,w)}G()}b.value++}const B=!vt();let E=!1;function J(i){var c;(c=e.onScroll)===null||c===void 0||c.call(e,i),(!B||!E)&&G()}function Q(i){var c;if((c=e.onWheel)===null||c===void 0||c.call(e,i),B){const g=a.value;if(g!=null){if(i.deltaX===0&&(g.scrollTop===0&&i.deltaY<=0||g.scrollTop+g.offsetHeight>=g.scrollHeight&&i.deltaY>=0))return;i.preventDefault(),g.scrollTop+=i.deltaY/ge(),g.scrollLeft+=i.deltaX/ge(),G(),E=!0,Ge(()=>{E=!1})}}}function ee(i){if(o||U(i.target)||i.contentRect.height===u.value)return;u.value=i.contentRect.height;const{onResize:c}=e;c!==void 0&&c(i)}function G(){const{value:i}=a;i!=null&&(m.value=i.scrollTop,f=i.scrollLeft)}function U(i){let c=i;for(;c!==null;){if(c.style.display==="none")return!0;c=c.parentElement}return!1}return{listHeight:u,listStyle:{overflow:"auto"},keyToIndex:r,itemsStyle:T(()=>{const{itemResizable:i}=e,c=K(R.value.sum());return b.value,[e.itemsStyle,{boxSizing:"content-box",height:i?"":c,minHeight:i?c:"",paddingTop:K(e.paddingTop),paddingBottom:K(e.paddingBottom)}]}),visibleItemsStyle:T(()=>(b.value,{transform:`translateY(${K(R.value.sum(z.value))})`})),viewportItems:F,listElRef:a,itemsElRef:$(null),scrollTo:h,handleListResize:ee,handleListScroll:J,handleListWheel:Q,handleItemResize:I}},render(){const{itemResizable:e,keyField:t,keyToIndex:o,visibleItemsTag:s}=this;return d(de,{onResize:this.handleListResize},{default:()=>{var r,a;return d("div",ye(this.$attrs,{class:["v-vl",this.showScrollbar&&"v-vl--show-scrollbar"],onScroll:this.handleListScroll,onWheel:this.handleListWheel,ref:"listElRef"}),[this.items.length!==0?d("div",{ref:"itemsElRef",class:"v-vl-items",style:this.itemsStyle},[d(s,Object.assign({class:"v-vl-visible-items",style:this.visibleItemsStyle},this.visibleItemsProps),{default:()=>this.viewportItems.map(u=>{const v=u[t],R=o.get(v),b=this.$slots.default({item:u,index:R})[0];return e?d(de,{key:v,onResize:f=>this.handleItemResize(v,f)},{default:()=>b}):(b.key=v,b)})})]):(a=(r=this.$slots).empty)===null||a===void 0?void 0:a.call(r)])}})}});function gt(e,t){t&&(se(()=>{const{value:o}=e;o&&ce.registerHandler(o,t)}),xe(()=>{const{value:o}=e;o&&ce.unregisterHandler(o)}))}var pt=O({name:"Checkmark",render(){return d("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 16 16"},d("g",{fill:"none"},d("path",{d:"M14.046 3.486a.75.75 0 0 1-.032 1.06l-7.93 7.474a.85.85 0 0 1-1.188-.022l-2.68-2.72a.75.75 0 1 1 1.068-1.053l2.234 2.267l7.468-7.038a.75.75 0 0 1 1.06.032z",fill:"currentColor"})))}}),bt=O({name:"Empty",render(){return d("svg",{viewBox:"0 0 28 28",fill:"none",xmlns:"http://www.w3.org/2000/svg"},d("path",{d:"M26 7.5C26 11.0899 23.0899 14 19.5 14C15.9101 14 13 11.0899 13 7.5C13 3.91015 15.9101 1 19.5 1C23.0899 1 26 3.91015 26 7.5ZM16.8536 4.14645C16.6583 3.95118 16.3417 3.95118 16.1464 4.14645C15.9512 4.34171 15.9512 4.65829 16.1464 4.85355L18.7929 7.5L16.1464 10.1464C15.9512 10.3417 15.9512 10.6583 16.1464 10.8536C16.3417 11.0488 16.6583 11.0488 16.8536 10.8536L19.5 8.20711L22.1464 10.8536C22.3417 11.0488 22.6583 11.0488 22.8536 10.8536C23.0488 10.6583 23.0488 10.3417 22.8536 10.1464L20.2071 7.5L22.8536 4.85355C23.0488 4.65829 23.0488 4.34171 22.8536 4.14645C22.6583 3.95118 22.3417 3.95118 22.1464 4.14645L19.5 6.79289L16.8536 4.14645Z",fill:"currentColor"}),d("path",{d:"M25 22.75V12.5991C24.5572 13.0765 24.053 13.4961 23.5 13.8454V16H17.5L17.3982 16.0068C17.0322 16.0565 16.75 16.3703 16.75 16.75C16.75 18.2688 15.5188 19.5 14 19.5C12.4812 19.5 11.25 18.2688 11.25 16.75L11.2432 16.6482C11.1935 16.2822 10.8797 16 10.5 16H4.5V7.25C4.5 6.2835 5.2835 5.5 6.25 5.5H12.2696C12.4146 4.97463 12.6153 4.47237 12.865 4H6.25C4.45507 4 3 5.45507 3 7.25V22.75C3 24.5449 4.45507 26 6.25 26H21.75C23.5449 26 25 24.5449 25 22.75ZM4.5 22.75V17.5H9.81597L9.85751 17.7041C10.2905 19.5919 11.9808 21 14 21L14.215 20.9947C16.2095 20.8953 17.842 19.4209 18.184 17.5H23.5V22.75C23.5 23.7165 22.7165 24.5 21.75 24.5H6.25C5.2835 24.5 4.5 23.7165 4.5 22.75Z",fill:"currentColor"}))}}),yt=O({props:{onFocus:Function,onBlur:Function},setup(e){return()=>d("div",{style:"width: 0; height: 0",tabindex:0,onFocus:e.onFocus,onBlur:e.onBlur})}}),xt=L("empty",`
 display: flex;
 flex-direction: column;
 align-items: center;
 font-size: var(--n-font-size);
`,[V("icon",`
 width: var(--n-icon-size);
 height: var(--n-icon-size);
 font-size: var(--n-icon-size);
 line-height: var(--n-icon-size);
 color: var(--n-icon-color);
 transition:
 color .3s var(--n-bezier);
 `,[j("+",[V("description",`
 margin-top: 8px;
 `)])]),V("description",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 `),V("extra",`
 text-align: center;
 transition: color .3s var(--n-bezier);
 margin-top: 12px;
 color: var(--n-extra-text-color);
 `)]);const wt=Object.assign(Object.assign({},W.props),{description:String,showDescription:{type:Boolean,default:!0},showIcon:{type:Boolean,default:!0},size:{type:String,default:"medium"},renderIcon:Function});var St=O({name:"Empty",props:wt,setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:o}=we(e),s=W("Empty","-empty",xt,qe,e,t),{localeRef:r}=Ue("Empty"),a=Z(Xe,null),u=T(()=>{var f,m,z;return(f=e.description)!==null&&f!==void 0?f:(z=(m=a==null?void 0:a.mergedComponentPropsRef.value)===null||m===void 0?void 0:m.Empty)===null||z===void 0?void 0:z.description}),v=T(()=>{var f,m;return((m=(f=a==null?void 0:a.mergedComponentPropsRef.value)===null||f===void 0?void 0:f.Empty)===null||m===void 0?void 0:m.renderIcon)||(()=>d(bt,null))}),R=T(()=>{const{size:f}=e,{common:{cubicBezierEaseInOut:m},self:{[D("iconSize",f)]:z,[D("fontSize",f)]:F,textColor:h,iconColor:y,extraTextColor:M}}=s.value;return{"--n-icon-size":z,"--n-font-size":F,"--n-bezier":m,"--n-text-color":h,"--n-icon-color":y,"--n-extra-text-color":M}}),b=o?Se("empty",T(()=>{let f="";const{size:m}=e;return f+=m[0],f}),R,e):void 0;return{mergedClsPrefix:t,mergedRenderIcon:v,localizedDescription:T(()=>u.value||r.value.description),cssVars:o?void 0:R,themeClass:b==null?void 0:b.themeClass,onRender:b==null?void 0:b.onRender}},render(){const{$slots:e,mergedClsPrefix:t,onRender:o}=this;return o==null||o(),d("div",{class:[`${t}-empty`,this.themeClass],style:this.cssVars},this.showIcon?d("div",{class:`${t}-empty__icon`},e.icon?e.icon():d(ze,{clsPrefix:t},{default:this.mergedRenderIcon})):null,this.showDescription?d("div",{class:`${t}-empty__description`},e.default?e.default():this.localizedDescription):null,e.extra?d("div",{class:`${t}-empty__extra`},e.extra()):null)}});function zt(e,t){return d(Ye,{name:"fade-in-scale-up-transition"},{default:()=>e?d(ze,{clsPrefix:t,class:`${t}-base-select-option__check`},{default:()=>d(pt)}):null})}var pe=O({name:"NBaseSelectOption",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(e){const{valueRef:t,pendingTmNodeRef:o,multipleRef:s,valueSetRef:r,renderLabelRef:a,renderOptionRef:u,labelFieldRef:v,valueFieldRef:R,showCheckmarkRef:b,nodePropsRef:f,handleOptionClick:m,handleOptionMouseEnter:z}=Z(ae),F=Y(()=>{const{value:S}=o;return S?e.tmNode.key===S.key:!1});function h(S){const{tmNode:x}=e;x.disabled||m(S,x)}function y(S){const{tmNode:x}=e;x.disabled||z(S,x)}function M(S){const{tmNode:x}=e,{value:I}=F;x.disabled||I||z(S,x)}return{multiple:s,isGrouped:Y(()=>{const{tmNode:S}=e,{parent:x}=S;return x&&x.rawNode.type==="group"}),showCheckmark:b,nodeProps:f,isPending:F,isSelected:Y(()=>{const{value:S}=t,{value:x}=s;if(S===null)return!1;const I=e.tmNode.rawNode[R.value];if(x){const{value:B}=r;return B.has(I)}else return S===I}),labelField:v,renderLabel:a,renderOption:u,handleMouseMove:M,handleMouseEnter:y,handleClick:h}},render(){const{clsPrefix:e,tmNode:{rawNode:t},isSelected:o,isPending:s,isGrouped:r,showCheckmark:a,nodeProps:u,renderOption:v,renderLabel:R,handleClick:b,handleMouseEnter:f,handleMouseMove:m}=this,z=zt(o,e),F=R?[R(t,o),a&&z]:[Re(t[this.labelField],t,o),a&&z],h=u==null?void 0:u(t),y=d("div",Object.assign({},h,{class:[`${e}-base-select-option`,t.class,h==null?void 0:h.class,{[`${e}-base-select-option--disabled`]:t.disabled,[`${e}-base-select-option--selected`]:o,[`${e}-base-select-option--grouped`]:r,[`${e}-base-select-option--pending`]:s,[`${e}-base-select-option--show-checkmark`]:a}],style:[(h==null?void 0:h.style)||"",t.style||""],onClick:ie([b,h==null?void 0:h.onClick]),onMouseenter:ie([f,h==null?void 0:h.onMouseenter]),onMousemove:ie([m,h==null?void 0:h.onMousemove])}),d("div",{class:`${e}-base-select-option__content`},F));return t.render?t.render({node:y,option:t,selected:o}):v?v({node:y,option:t,selected:o}):y}}),be=O({name:"NBaseSelectGroupHeader",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(){const{renderLabelRef:e,renderOptionRef:t,labelFieldRef:o,nodePropsRef:s}=Z(ae);return{labelField:o,nodeProps:s,renderLabel:e,renderOption:t}},render(){const{clsPrefix:e,renderLabel:t,renderOption:o,nodeProps:s,tmNode:{rawNode:r}}=this,a=s==null?void 0:s(r),u=t?t(r,!1):Re(r[this.labelField],r,!1),v=d("div",Object.assign({},a,{class:[`${e}-base-select-group-header`,a==null?void 0:a.class]}),u);return r.render?r.render({node:v,option:r}):o?o({node:v,option:r,selected:!1}):v}}),Rt=L("base-select-menu",`
 line-height: 1.5;
 outline: none;
 z-index: 0;
 position: relative;
 border-radius: var(--n-border-radius);
 transition:
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-color);
`,[L("scrollbar",`
 max-height: var(--n-height);
 `),L("virtual-list",`
 max-height: var(--n-height);
 `),L("base-select-option",`
 min-height: var(--n-option-height);
 font-size: var(--n-option-font-size);
 display: flex;
 align-items: center;
 `,[V("content",`
 z-index: 1;
 white-space: nowrap;
 text-overflow: ellipsis;
 overflow: hidden;
 `)]),L("base-select-group-header",`
 min-height: var(--n-option-height);
 font-size: .93em;
 display: flex;
 align-items: center;
 `),L("base-select-menu-option-wrapper",`
 position: relative;
 width: 100%;
 `),V("loading, empty",`
 display: flex;
 padding: 12px 32px;
 flex: 1;
 justify-content: center;
 `),V("loading",`
 color: var(--n-loading-color);
 font-size: var(--n-loading-size);
 `),V("action",`
 padding: 8px var(--n-option-padding-left);
 font-size: var(--n-option-font-size);
 transition: 
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 border-top: 1px solid var(--n-action-divider-color);
 color: var(--n-action-text-color);
 `),L("base-select-group-header",`
 position: relative;
 cursor: default;
 padding: var(--n-option-padding);
 color: var(--n-group-header-text-color);
 `),L("base-select-option",`
 cursor: pointer;
 position: relative;
 padding: var(--n-option-padding);
 transition:
 color .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 box-sizing: border-box;
 color: var(--n-option-text-color);
 opacity: 1;
 `,[A("show-checkmark",`
 padding-right: calc(var(--n-option-padding-right) + 20px);
 `),j("&::before",`
 content: "";
 position: absolute;
 left: 4px;
 right: 4px;
 top: 0;
 bottom: 0;
 border-radius: var(--n-border-radius);
 transition: background-color .3s var(--n-bezier);
 `),j("&:active",`
 color: var(--n-option-text-color-pressed);
 `),A("grouped",`
 padding-left: calc(var(--n-option-padding-left) * 1.5);
 `),A("pending",[j("&::before",`
 background-color: var(--n-option-color-pending);
 `)]),A("selected",`
 color: var(--n-option-text-color-active);
 `,[j("&::before",`
 background-color: var(--n-option-color-active);
 `),A("pending",[j("&::before",`
 background-color: var(--n-option-color-active-pending);
 `)])]),A("disabled",`
 cursor: not-allowed;
 `,[Ze("selected",`
 color: var(--n-option-text-color-disabled);
 `),A("selected",`
 opacity: var(--n-option-opacity-disabled);
 `)]),V("check",`
 font-size: 16px;
 position: absolute;
 right: calc(var(--n-option-padding-right) - 4px);
 top: calc(50% - 7px);
 color: var(--n-option-check-color);
 transition: color .3s var(--n-bezier);
 `,[Je({enterScale:"0.5"})])])]),Nt=O({name:"InternalSelectMenu",props:Object.assign(Object.assign({},W.props),{clsPrefix:{type:String,required:!0},scrollable:{type:Boolean,default:!0},treeMate:{type:Object,required:!0},multiple:Boolean,size:{type:String,default:"medium"},value:{type:[String,Number,Array],default:null},autoPending:Boolean,virtualScroll:{type:Boolean,default:!0},show:{type:Boolean,default:!0},labelField:{type:String,default:"label"},valueField:{type:String,default:"value"},loading:Boolean,focusable:Boolean,renderLabel:Function,renderOption:Function,nodeProps:Function,showCheckmark:{type:Boolean,default:!0},onMousedown:Function,onScroll:Function,onFocus:Function,onBlur:Function,onKeyup:Function,onKeydown:Function,onTabOut:Function,onMouseenter:Function,onMouseleave:Function,onResize:Function,resetMenuOnOptionsChange:{type:Boolean,default:!0},inlineThemeDisabled:Boolean,onToggle:Function}),setup(e){const t=W("InternalSelectMenu","-internal-select-menu",Rt,Qe,e,H(e,"clsPrefix")),o=$(null),s=$(null),r=$(null),a=T(()=>e.treeMate.getFlattenedNodes()),u=T(()=>et(a.value)),v=$(null);function R(){const{treeMate:n}=e;let l=null;const{value:w}=e;w===null?l=n.getFirstAvailableNode():(e.multiple?l=n.getNode((w||[])[(w||[]).length-1]):l=n.getNode(w),(!l||l.disabled)&&(l=n.getFirstAvailableNode())),i(l||null)}function b(){const{value:n}=v;n&&!e.treeMate.getNode(n.key)&&(v.value=null)}let f;ue(()=>e.show,n=>{n?f=ue(()=>e.treeMate,()=>{e.resetMenuOnOptionsChange?(e.autoPending?R():b(),tt(c)):b()},{immediate:!0}):f==null||f()},{immediate:!0}),xe(()=>{f==null||f()});const m=T(()=>re(t.value.self[D("optionHeight",e.size)])),z=T(()=>oe(t.value.self[D("padding",e.size)])),F=T(()=>e.multiple&&Array.isArray(e.value)?new Set(e.value):new Set),h=T(()=>{const n=a.value;return n&&n.length===0});function y(n){const{onToggle:l}=e;l&&l(n)}function M(n){const{onScroll:l}=e;l&&l(n)}function S(n){var l;(l=r.value)===null||l===void 0||l.sync(),M(n)}function x(){var n;(n=r.value)===null||n===void 0||n.sync()}function I(){const{value:n}=v;return n||null}function B(n,l){l.disabled||i(l,!1)}function E(n,l){l.disabled||y(l)}function J(n){var l;ve(n,"action")||(l=e.onKeyup)===null||l===void 0||l.call(e,n)}function Q(n){var l;ve(n,"action")||(l=e.onKeydown)===null||l===void 0||l.call(e,n)}function ee(n){var l;(l=e.onMousedown)===null||l===void 0||l.call(e,n),!e.focusable&&n.preventDefault()}function G(){const{value:n}=v;n&&i(n.getNext({loop:!0}),!0)}function U(){const{value:n}=v;n&&i(n.getPrev({loop:!0}),!0)}function i(n,l=!1){v.value=n,l&&c()}function c(){var n,l;const w=v.value;if(!w)return;const N=u.value(w.key);N!==null&&(e.virtualScroll?(n=s.value)===null||n===void 0||n.scrollTo({index:N}):(l=r.value)===null||l===void 0||l.scrollTo({index:N,elSize:m.value}))}function g(n){var l,w;!((l=o.value)===null||l===void 0)&&l.contains(n.target)&&((w=e.onFocus)===null||w===void 0||w.call(e,n))}function k(n){var l,w;!((l=o.value)===null||l===void 0)&&l.contains(n.relatedTarget)||(w=e.onBlur)===null||w===void 0||w.call(e,n)}fe(ae,{handleOptionMouseEnter:B,handleOptionClick:E,valueSetRef:F,pendingTmNodeRef:v,nodePropsRef:H(e,"nodeProps"),showCheckmarkRef:H(e,"showCheckmark"),multipleRef:H(e,"multiple"),valueRef:H(e,"value"),renderLabelRef:H(e,"renderLabel"),renderOptionRef:H(e,"renderOption"),labelFieldRef:H(e,"labelField"),valueFieldRef:H(e,"valueField")}),fe(nt,o),se(()=>{const{value:n}=r;n&&n.sync()});const C=T(()=>{const{size:n}=e,{common:{cubicBezierEaseInOut:l},self:{height:w,borderRadius:N,color:q,groupHeaderTextColor:ke,actionDividerColor:Ce,optionTextColorPressed:Te,optionTextColor:Pe,optionTextColorDisabled:_e,optionTextColorActive:Me,optionOpacityDisabled:Ne,optionCheckColor:Fe,actionTextColor:Ie,optionColorPending:Be,optionColorActive:Le,loadingColor:Oe,loadingSize:Ee,optionColorActivePending:He,[D("optionFontSize",n)]:$e,[D("optionHeight",n)]:Ve,[D("optionPadding",n)]:te}}=t.value;return{"--n-height":w,"--n-action-divider-color":Ce,"--n-action-text-color":Ie,"--n-bezier":l,"--n-border-radius":N,"--n-color":q,"--n-option-font-size":$e,"--n-group-header-text-color":ke,"--n-option-check-color":Fe,"--n-option-color-pending":Be,"--n-option-color-active":Le,"--n-option-color-active-pending":He,"--n-option-height":Ve,"--n-option-opacity-disabled":Ne,"--n-option-text-color":Pe,"--n-option-text-color-active":Me,"--n-option-text-color-disabled":_e,"--n-option-text-color-pressed":Te,"--n-option-padding":te,"--n-option-padding-left":oe(te,"left"),"--n-option-padding-right":oe(te,"right"),"--n-loading-color":Oe,"--n-loading-size":Ee}}),{inlineThemeDisabled:P}=e,p=P?Se("internal-select-menu",T(()=>e.size[0]),C,e):void 0,_={selfRef:o,next:G,prev:U,getPendingTmNode:I};return gt(o,e.onResize),Object.assign({mergedTheme:t,virtualListRef:s,scrollbarRef:r,itemSize:m,padding:z,flattenedNodes:a,empty:h,virtualListContainer(){const{value:n}=s;return n==null?void 0:n.listElRef},virtualListContent(){const{value:n}=s;return n==null?void 0:n.itemsElRef},doScroll:M,handleFocusin:g,handleFocusout:k,handleKeyUp:J,handleKeyDown:Q,handleMouseDown:ee,handleVirtualListResize:x,handleVirtualListScroll:S,cssVars:P?void 0:C,themeClass:p==null?void 0:p.themeClass,onRender:p==null?void 0:p.onRender},_)},render(){const{$slots:e,virtualScroll:t,clsPrefix:o,mergedTheme:s,themeClass:r,onRender:a}=this;return a==null||a(),d("div",{ref:"selfRef",tabindex:this.focusable?0:-1,class:[`${o}-base-select-menu`,r,this.multiple&&`${o}-base-select-menu--multiple`],style:this.cssVars,onFocusin:this.handleFocusin,onFocusout:this.handleFocusout,onKeyup:this.handleKeyUp,onKeydown:this.handleKeyDown,onMousedown:this.handleMouseDown,onMouseenter:this.onMouseenter,onMouseleave:this.onMouseleave},this.loading?d("div",{class:`${o}-base-select-menu__loading`},d(it,{clsPrefix:o,strokeWidth:20})):this.empty?d("div",{class:`${o}-base-select-menu__empty`,"data-empty":!0},rt(e.empty,()=>[d(St,{theme:s.peers.Empty,themeOverrides:s.peerOverrides.Empty})])):d(lt,{ref:"scrollbarRef",theme:s.peers.Scrollbar,themeOverrides:s.peerOverrides.Scrollbar,scrollable:this.scrollable,container:t?this.virtualListContainer:void 0,content:t?this.virtualListContent:void 0,onScroll:t?void 0:this.doScroll},{default:()=>t?d(mt,{ref:"virtualListRef",class:`${o}-virtual-list`,items:this.flattenedNodes,itemSize:this.itemSize,showScrollbar:!1,paddingTop:this.padding.top,paddingBottom:this.padding.bottom,onResize:this.handleVirtualListResize,onScroll:this.handleVirtualListScroll,itemResizable:!0},{default:({item:u})=>u.isGroup?d(be,{key:u.key,clsPrefix:o,tmNode:u}):u.ignored?null:d(pe,{clsPrefix:o,key:u.key,tmNode:u})}):d("div",{class:`${o}-base-select-menu-option-wrapper`,style:{paddingTop:this.padding.top,paddingBottom:this.padding.bottom}},this.flattenedNodes.map(u=>u.isGroup?d(be,{key:u.key,clsPrefix:o,tmNode:u}):d(pe,{clsPrefix:o,key:u.key,tmNode:u})))}),ot(e.action,u=>u&&[d("div",{class:`${o}-base-select-menu__action`,"data-action":!0,key:"action"},u),d(yt,{onFocus:this.onTabOut,key:"focus-detector"})]))}}),Ft=O({name:"ListItem",setup(){const e=Z(ct,null);return e||st("list-item","`n-list-item` must be placed in `n-list`."),{showDivider:e.showDividerRef,mergedClsPrefix:e.mergedClsPrefixRef}},render(){const{$slots:e,mergedClsPrefix:t}=this;return d("li",{class:`${t}-list-item`},e.prefix?d("div",{class:`${t}-list-item__prefix`},e.prefix()):null,e.default?d("div",{class:`${t}-list-item__main`},e):null,e.suffix?d("div",{class:`${t}-list-item__suffix`},e.suffix()):null,this.showDivider&&d("div",{class:`${t}-list-item__divider`}))}});const kt=e=>{const{heightSmall:t,heightMedium:o,heightLarge:s,borderRadius:r}=e;return{color:"#eee",colorEnd:"#ddd",borderRadius:r,heightSmall:t,heightMedium:o,heightLarge:s}},Ct={name:"Skeleton",common:at,self:kt};var Tt=j([L("skeleton",`
 height: 1em;
 width: 100%;
 transition: background-color .3s var(--n-bezier);
 transition:
 --n-color-start .3s var(--n-bezier),
 --n-color-end .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 animation: 2s skeleton-loading infinite cubic-bezier(0.36, 0, 0.64, 1);
 background-color: var(--n-color-start);
 `),j("@keyframes skeleton-loading",`
 0% {
 background: var(--n-color-start);
 }
 40% {
 background: var(--n-color-end);
 }
 80% {
 background: var(--n-color-start);
 }
 100% {
 background: var(--n-color-start);
 }
 `)]);const Pt=Object.assign(Object.assign({},W.props),{text:Boolean,round:Boolean,circle:Boolean,height:[String,Number],width:[String,Number],size:String,repeat:{type:Number,default:1},animated:{type:Boolean,default:!0},sharp:{type:Boolean,default:!0}});var It=O({name:"Skeleton",inheritAttrs:!1,props:Pt,setup(e){ut();const{mergedClsPrefixRef:t}=we(e),o=W("Skeleton","-skeleton",Tt,Ct,e,t);return{mergedClsPrefix:t,style:T(()=>{var s,r;const a=o.value,{common:{cubicBezierEaseInOut:u}}=a,v=a.self,{color:R,colorEnd:b,borderRadius:f}=v;let m;const{circle:z,sharp:F,round:h,width:y,height:M,size:S,text:x,animated:I}=e;S!==void 0&&(m=v[D("height",S)]);const B=z?(s=y!=null?y:M)!==null&&s!==void 0?s:m:y,E=(r=z&&y!=null?y:M)!==null&&r!==void 0?r:m;return{display:x?"inline-block":"",verticalAlign:x?"-0.125em":"",borderRadius:z?"50%":h?"4096px":F?"":f,width:typeof B=="number"?K(B):B,height:typeof E=="number"?K(E):E,animation:I?"":"none","--n-bezier":u,"--n-color-start":R,"--n-color-end":b}})}},render(){const{repeat:e,style:t,mergedClsPrefix:o,$attrs:s}=this,r=d("div",ye({class:`${o}-skeleton`,style:t},s));return e>1?d(dt,null,Array.apply(null,{length:e}).map(a=>[r,`
`])):r}});export{Nt as N,Ft as _,St as a,It as b,ut as c,ie as m,gt as u};
