import{d as le,r as P,bQ as Zt,bR as Qt,a2 as kt,V as Ue,h as n,bS as Yt,bT as Xt,b as ce,c as k,f as U,a as it,e as G,x as Ce,t as we,bU as Gt,y as I,bV as We,S as Ne,bJ as He,z as Q,A as Ze,bW as en,bX as tn,aL as at,as as Pt,ab as lt,bY as nn,n as on,q as an,u as dt,bZ as Ot,aW as Bt,b_ as st,w as E,ao as rn,aq as ln,b$ as sn,ar as zt,at as ct,p as dn,aV as un,c0 as cn,s as Ke,J as Rt,c1 as fn,o as hn,aY as vn,aX as qe,aZ as bn,a_ as gn,a$ as pn,b0 as mn,bw as wn,bo as Cn,c2 as ft,c3 as xn,c4 as yn,c5 as Fn,i as Mn,L as Sn,_ as ht,N as Re}from"./index-476ef983.js";import{c as kn,N as Tt,m as vt}from"./Skeleton-3d68b5c2.js";function bt(e){switch(e){case"tiny":return"mini";case"small":return"tiny";case"medium":return"small";case"large":return"medium";case"huge":return"large"}throw Error(`${e} has no smaller size.`)}const ye="v-hidden",Pn=Xt("[v-hidden]",{display:"none!important"}),gt=le({name:"Overflow",props:{getCounter:Function,getTail:Function,updateCounter:Function,onUpdateOverflow:Function},setup(e,{slots:o}){const s=P(null),d=P(null);function f(){const{value:C}=s,{getCounter:i,getTail:b}=e;let h;if(i!==void 0?h=i():h=d.value,!C||!h)return;h.hasAttribute(ye)&&h.removeAttribute(ye);const{children:F}=C,y=C.offsetWidth,v=[],g=o.tail?b==null?void 0:b():null;let c=g?g.offsetWidth:0,p=!1;const z=C.children.length-(o.tail?1:0);for(let R=0;R<z-1;++R){if(R<0)continue;const W=F[R];if(p){W.hasAttribute(ye)||W.setAttribute(ye,"");continue}else W.hasAttribute(ye)&&W.removeAttribute(ye);const Y=W.offsetWidth;if(c+=Y,v[R]=Y,c>y){const{updateCounter:H}=e;for(let _=R;_>=0;--_){const L=z-1-_;H!==void 0?H(L):h.textContent=`${L}`;const N=h.offsetWidth;if(c-=v[_],c+N<=y||_===0){p=!0,R=_-1,g&&(R===-1?(g.style.maxWidth=`${y-N}px`,g.style.boxSizing="border-box"):g.style.maxWidth="");break}}}}const{onUpdateOverflow:T}=e;p?T!==void 0&&T(!0):(T!==void 0&&T(!1),h.setAttribute(ye,""))}const x=Zt();return Pn.mount({id:"vueuc/overflow",head:!0,anchorMetaName:Qt,ssr:x}),kt(f),{selfRef:s,counterRef:d,sync:f}},render(){const{$slots:e}=this;return Ue(this.sync),n("div",{class:"v-overflow",ref:"selfRef"},[Yt(e,"default"),e.counter?e.counter():n("span",{style:{display:"inline-block"},ref:"counterRef"}),e.tail?e.tail():null])}}),pt=le({name:"Backward",render(){return n("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},n("path",{d:"M12.2674 15.793C11.9675 16.0787 11.4927 16.0672 11.2071 15.7673L6.20572 10.5168C5.9298 10.2271 5.9298 9.7719 6.20572 9.48223L11.2071 4.23177C11.4927 3.93184 11.9675 3.92031 12.2674 4.206C12.5673 4.49169 12.5789 4.96642 12.2932 5.26634L7.78458 9.99952L12.2932 14.7327C12.5789 15.0326 12.5673 15.5074 12.2674 15.793Z",fill:"currentColor"}))}}),mt=le({name:"FastBackward",render(){return n("svg",{viewBox:"0 0 20 20",version:"1.1",xmlns:"http://www.w3.org/2000/svg"},n("g",{stroke:"none","stroke-width":"1",fill:"none","fill-rule":"evenodd"},n("g",{fill:"currentColor","fill-rule":"nonzero"},n("path",{d:"M8.73171,16.7949 C9.03264,17.0795 9.50733,17.0663 9.79196,16.7654 C10.0766,16.4644 10.0634,15.9897 9.76243,15.7051 L4.52339,10.75 L17.2471,10.75 C17.6613,10.75 17.9971,10.4142 17.9971,10 C17.9971,9.58579 17.6613,9.25 17.2471,9.25 L4.52112,9.25 L9.76243,4.29275 C10.0634,4.00812 10.0766,3.53343 9.79196,3.2325 C9.50733,2.93156 9.03264,2.91834 8.73171,3.20297 L2.31449,9.27241 C2.14819,9.4297 2.04819,9.62981 2.01448,9.8386 C2.00308,9.89058 1.99707,9.94459 1.99707,10 C1.99707,10.0576 2.00356,10.1137 2.01585,10.1675 C2.05084,10.3733 2.15039,10.5702 2.31449,10.7254 L8.73171,16.7949 Z"}))))}}),wt=le({name:"FastForward",render(){return n("svg",{viewBox:"0 0 20 20",version:"1.1",xmlns:"http://www.w3.org/2000/svg"},n("g",{stroke:"none","stroke-width":"1",fill:"none","fill-rule":"evenodd"},n("g",{fill:"currentColor","fill-rule":"nonzero"},n("path",{d:"M11.2654,3.20511 C10.9644,2.92049 10.4897,2.93371 10.2051,3.23464 C9.92049,3.53558 9.93371,4.01027 10.2346,4.29489 L15.4737,9.25 L2.75,9.25 C2.33579,9.25 2,9.58579 2,10.0000012 C2,10.4142 2.33579,10.75 2.75,10.75 L15.476,10.75 L10.2346,15.7073 C9.93371,15.9919 9.92049,16.4666 10.2051,16.7675 C10.4897,17.0684 10.9644,17.0817 11.2654,16.797 L17.6826,10.7276 C17.8489,10.5703 17.9489,10.3702 17.9826,10.1614 C17.994,10.1094 18,10.0554 18,10.0000012 C18,9.94241 17.9935,9.88633 17.9812,9.83246 C17.9462,9.62667 17.8467,9.42976 17.6826,9.27455 L11.2654,3.20511 Z"}))))}}),Ct=le({name:"Forward",render(){return n("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},n("path",{d:"M7.73271 4.20694C8.03263 3.92125 8.50737 3.93279 8.79306 4.23271L13.7944 9.48318C14.0703 9.77285 14.0703 10.2281 13.7944 10.5178L8.79306 15.7682C8.50737 16.0681 8.03263 16.0797 7.73271 15.794C7.43279 15.5083 7.42125 15.0336 7.70694 14.7336L12.2155 10.0005L7.70694 5.26729C7.42125 4.96737 7.43279 4.49264 7.73271 4.20694Z",fill:"currentColor"}))}}),xt=le({name:"More",render(){return n("svg",{viewBox:"0 0 16 16",version:"1.1",xmlns:"http://www.w3.org/2000/svg"},n("g",{stroke:"none","stroke-width":"1",fill:"none","fill-rule":"evenodd"},n("g",{fill:"currentColor","fill-rule":"nonzero"},n("path",{d:"M4,7 C4.55228,7 5,7.44772 5,8 C5,8.55229 4.55228,9 4,9 C3.44772,9 3,8.55229 3,8 C3,7.44772 3.44772,7 4,7 Z M8,7 C8.55229,7 9,7.44772 9,8 C9,8.55229 8.55229,9 8,9 C7.44772,9 7,8.55229 7,8 C7,7.44772 7.44772,7 8,7 Z M12,7 C12.5523,7 13,7.44772 13,8 C13,8.55229 12.5523,9 12,9 C11.4477,9 11,8.55229 11,8 C11,7.44772 11.4477,7 12,7 Z"}))))}}),On=ce([k("base-selection",`
 position: relative;
 z-index: auto;
 box-shadow: none;
 width: 100%;
 max-width: 100%;
 display: inline-block;
 vertical-align: bottom;
 border-radius: var(--n-border-radius);
 min-height: var(--n-height);
 line-height: 1.5;
 font-size: var(--n-font-size);
 `,[k("base-loading",`
 color: var(--n-loading-color);
 `),k("base-selection-tags","min-height: var(--n-height);"),U("border, state-border",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border: var(--n-border);
 border-radius: inherit;
 transition:
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `),U("state-border",`
 z-index: 1;
 border-color: #0000;
 `),k("base-suffix",`
 cursor: pointer;
 position: absolute;
 top: 50%;
 transform: translateY(-50%);
 right: 10px;
 `,[U("arrow",`
 font-size: var(--n-arrow-size);
 color: var(--n-arrow-color);
 transition: color .3s var(--n-bezier);
 `)]),k("base-selection-overlay",`
 display: flex;
 align-items: center;
 white-space: nowrap;
 pointer-events: none;
 position: absolute;
 top: 0;
 right: 0;
 bottom: 0;
 left: 0;
 padding: var(--n-padding-single);
 transition: color .3s var(--n-bezier);
 `,[U("wrapper",`
 flex-basis: 0;
 flex-grow: 1;
 overflow: hidden;
 text-overflow: ellipsis;
 `)]),k("base-selection-placeholder",`
 color: var(--n-placeholder-color);
 `,[U("inner",`
 max-width: 100%;
 overflow: hidden;
 `)]),k("base-selection-tags",`
 cursor: pointer;
 outline: none;
 box-sizing: border-box;
 position: relative;
 z-index: auto;
 display: flex;
 padding: var(--n-padding-multiple);
 flex-wrap: wrap;
 align-items: center;
 width: 100%;
 vertical-align: bottom;
 background-color: var(--n-color);
 border-radius: inherit;
 transition:
 color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),k("base-selection-label",`
 height: var(--n-height);
 display: inline-flex;
 width: 100%;
 vertical-align: bottom;
 cursor: pointer;
 outline: none;
 z-index: auto;
 box-sizing: border-box;
 position: relative;
 transition:
 color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 border-radius: inherit;
 background-color: var(--n-color);
 align-items: center;
 `,[k("base-selection-input",`
 font-size: inherit;
 line-height: inherit;
 outline: none;
 cursor: pointer;
 box-sizing: border-box;
 border:none;
 width: 100%;
 padding: var(--n-padding-single);
 background-color: #0000;
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 caret-color: var(--n-caret-color);
 `,[U("content",`
 text-overflow: ellipsis;
 overflow: hidden;
 white-space: nowrap; 
 `)]),U("render-label",`
 color: var(--n-text-color);
 `)]),it("disabled",[ce("&:hover",[U("state-border",`
 box-shadow: var(--n-box-shadow-hover);
 border: var(--n-border-hover);
 `)]),G("focus",[U("state-border",`
 box-shadow: var(--n-box-shadow-focus);
 border: var(--n-border-focus);
 `)]),G("active",[U("state-border",`
 box-shadow: var(--n-box-shadow-active);
 border: var(--n-border-active);
 `),k("base-selection-label","background-color: var(--n-color-active);"),k("base-selection-tags","background-color: var(--n-color-active);")])]),G("disabled","cursor: not-allowed;",[U("arrow",`
 color: var(--n-arrow-color-disabled);
 `),k("base-selection-label",`
 cursor: not-allowed;
 background-color: var(--n-color-disabled);
 `,[k("base-selection-input",`
 cursor: not-allowed;
 color: var(--n-text-color-disabled);
 `),U("render-label",`
 color: var(--n-text-color-disabled);
 `)]),k("base-selection-tags",`
 cursor: not-allowed;
 background-color: var(--n-color-disabled);
 `),k("base-selection-placeholder",`
 cursor: not-allowed;
 color: var(--n-placeholder-color-disabled);
 `)]),k("base-selection-input-tag",`
 height: calc(var(--n-height) - 6px);
 line-height: calc(var(--n-height) - 6px);
 outline: none;
 display: none;
 position: relative;
 margin-bottom: 3px;
 max-width: 100%;
 vertical-align: bottom;
 `,[U("input",`
 font-size: inherit;
 font-family: inherit;
 min-width: 1px;
 padding: 0;
 background-color: #0000;
 outline: none;
 border: none;
 max-width: 100%;
 overflow: hidden;
 width: 1em;
 line-height: inherit;
 cursor: pointer;
 color: var(--n-text-color);
 caret-color: var(--n-caret-color);
 `),U("mirror",`
 position: absolute;
 left: 0;
 top: 0;
 white-space: pre;
 visibility: hidden;
 user-select: none;
 -webkit-user-select: none;
 opacity: 0;
 `)]),["warning","error"].map(e=>G(`${e}-status`,[U("state-border",`border: var(--n-border-${e});`),it("disabled",[ce("&:hover",[U("state-border",`
 box-shadow: var(--n-box-shadow-hover-${e});
 border: var(--n-border-hover-${e});
 `)]),G("active",[U("state-border",`
 box-shadow: var(--n-box-shadow-active-${e});
 border: var(--n-border-active-${e});
 `),k("base-selection-label",`background-color: var(--n-color-active-${e});`),k("base-selection-tags",`background-color: var(--n-color-active-${e});`)]),G("focus",[U("state-border",`
 box-shadow: var(--n-box-shadow-focus-${e});
 border: var(--n-border-focus-${e});
 `)])])]))]),k("base-selection-popover",`
 margin-bottom: -3px;
 display: flex;
 flex-wrap: wrap;
 margin-right: -8px;
 `),k("base-selection-tag-wrapper",`
 max-width: 100%;
 display: inline-flex;
 padding: 0 7px 3px 0;
 `,[ce("&:last-child","padding-right: 0;"),k("tag",`
 font-size: 14px;
 max-width: 100%;
 `,[U("content",`
 line-height: 1.25;
 text-overflow: ellipsis;
 overflow: hidden;
 `)])])]),Bn=le({name:"InternalSelection",props:Object.assign(Object.assign({},Ce.props),{clsPrefix:{type:String,required:!0},bordered:{type:Boolean,default:void 0},active:Boolean,pattern:{type:String,default:""},placeholder:String,selectedOption:{type:Object,default:null},selectedOptions:{type:Array,default:null},labelField:{type:String,default:"label"},valueField:{type:String,default:"value"},multiple:Boolean,filterable:Boolean,clearable:Boolean,disabled:Boolean,size:{type:String,default:"medium"},loading:Boolean,autofocus:Boolean,showArrow:{type:Boolean,default:!0},inputProps:Object,focused:Boolean,renderTag:Function,onKeydown:Function,onClick:Function,onBlur:Function,onFocus:Function,onDeleteOption:Function,maxTagCount:[String,Number],onClear:Function,onPatternInput:Function,onPatternFocus:Function,onPatternBlur:Function,renderLabel:Function,status:String,inlineThemeDisabled:Boolean,ignoreComposition:{type:Boolean,default:!0},onResize:Function}),setup(e){const o=P(null),s=P(null),d=P(null),f=P(null),x=P(null),C=P(null),i=P(null),b=P(null),h=P(null),F=P(null),y=P(!1),v=P(!1),g=P(!1),c=Ce("InternalSelection","-internal-selection",On,Gt,e,we(e,"clsPrefix")),p=I(()=>e.clearable&&!e.disabled&&(g.value||e.active)),z=I(()=>e.selectedOption?e.renderTag?e.renderTag({option:e.selectedOption,handleClose:()=>{}}):e.renderLabel?e.renderLabel(e.selectedOption,!0):We(e.selectedOption[e.labelField],e.selectedOption,!0):e.placeholder),T=I(()=>{const r=e.selectedOption;if(r)return r[e.labelField]}),R=I(()=>e.multiple?!!(Array.isArray(e.selectedOptions)&&e.selectedOptions.length):e.selectedOption!==null);function W(){var r;const{value:u}=o;if(u){const{value:$}=s;$&&($.style.width=`${u.offsetWidth}px`,e.maxTagCount!=="responsive"&&((r=h.value)===null||r===void 0||r.sync()))}}function Y(){const{value:r}=F;r&&(r.style.display="none")}function H(){const{value:r}=F;r&&(r.style.display="inline-block")}Ne(we(e,"active"),r=>{r||Y()}),Ne(we(e,"pattern"),()=>{e.multiple&&Ue(W)});function _(r){const{onFocus:u}=e;u&&u(r)}function L(r){const{onBlur:u}=e;u&&u(r)}function N(r){const{onDeleteOption:u}=e;u&&u(r)}function fe(r){const{onClear:u}=e;u&&u(r)}function ne(r){const{onPatternInput:u}=e;u&&u(r)}function se(r){var u;(!r.relatedTarget||!(!((u=d.value)===null||u===void 0)&&u.contains(r.relatedTarget)))&&_(r)}function V(r){var u;!((u=d.value)===null||u===void 0)&&u.contains(r.relatedTarget)||L(r)}function oe(r){fe(r)}function he(){g.value=!0}function ve(){g.value=!1}function J(r){!e.active||!e.filterable||r.target!==s.value&&r.preventDefault()}function K(r){N(r)}function X(r){if(r.key==="Backspace"&&!ae.value&&!e.pattern.length){const{selectedOptions:u}=e;u!=null&&u.length&&K(u[u.length-1])}}const ae=P(!1);let A=null;function xe(r){const{value:u}=o;if(u){const $=r.target.value;u.textContent=$,W()}e.ignoreComposition&&ae.value?A=r:ne(r)}function ee(){ae.value=!0}function be(){ae.value=!1,e.ignoreComposition&&ne(A),A=null}function ge(r){var u;v.value=!0,(u=e.onPatternFocus)===null||u===void 0||u.call(e,r)}function Z(r){var u;v.value=!1,(u=e.onPatternBlur)===null||u===void 0||u.call(e,r)}function pe(){var r,u;if(e.filterable)v.value=!1,(r=C.value)===null||r===void 0||r.blur(),(u=s.value)===null||u===void 0||u.blur();else if(e.multiple){const{value:$}=f;$==null||$.blur()}else{const{value:$}=x;$==null||$.blur()}}function re(){var r,u,$;e.filterable?(v.value=!1,(r=C.value)===null||r===void 0||r.focus()):e.multiple?(u=f.value)===null||u===void 0||u.focus():($=x.value)===null||$===void 0||$.focus()}function q(){const{value:r}=s;r&&(H(),r.focus())}function j(){const{value:r}=s;r&&r.blur()}function a(r){const{value:u}=i;u&&u.setTextContent(`+${r}`)}function m(){const{value:r}=b;return r}function ie(){return s.value}let te=null;function me(){te!==null&&window.clearTimeout(te)}function Te(){e.disabled||e.active||(me(),te=window.setTimeout(()=>{R.value&&(y.value=!0)},100))}function Ie(){me()}function _e(r){r||(me(),y.value=!1)}Ne(R,r=>{r||(y.value=!1)}),kt(()=>{He(()=>{const r=C.value;r&&(r.tabIndex=e.disabled||v.value?-1:0)})}),kn(d,e.onResize);const{inlineThemeDisabled:Fe}=e,Me=I(()=>{const{size:r}=e,{common:{cubicBezierEaseInOut:u},self:{borderRadius:$,color:Se,placeholderColor:$e,textColor:Ae,paddingSingle:je,paddingMultiple:Ee,caretColor:ke,colorDisabled:Pe,textColorDisabled:Oe,placeholderColorDisabled:Le,colorActive:Ve,boxShadowFocus:Be,boxShadowActive:ue,boxShadowHover:t,border:l,borderFocus:w,borderHover:B,borderActive:M,arrowColor:O,arrowColorDisabled:S,loadingColor:D,colorActiveWarning:ze,boxShadowFocusWarning:De,boxShadowActiveWarning:Qe,boxShadowHoverWarning:Ye,borderWarning:Xe,borderFocusWarning:Ge,borderHoverWarning:et,borderActiveWarning:tt,colorActiveError:nt,boxShadowFocusError:ot,boxShadowActiveError:At,boxShadowHoverError:jt,borderError:Et,borderFocusError:Lt,borderHoverError:Vt,borderActiveError:Dt,clearColor:Nt,clearColorHover:Ut,clearColorPressed:Wt,clearSize:Ht,arrowSize:Kt,[Q("height",r)]:qt,[Q("fontSize",r)]:Jt}}=c.value;return{"--n-bezier":u,"--n-border":l,"--n-border-active":M,"--n-border-focus":w,"--n-border-hover":B,"--n-border-radius":$,"--n-box-shadow-active":ue,"--n-box-shadow-focus":Be,"--n-box-shadow-hover":t,"--n-caret-color":ke,"--n-color":Se,"--n-color-active":Ve,"--n-color-disabled":Pe,"--n-font-size":Jt,"--n-height":qt,"--n-padding-single":je,"--n-padding-multiple":Ee,"--n-placeholder-color":$e,"--n-placeholder-color-disabled":Le,"--n-text-color":Ae,"--n-text-color-disabled":Oe,"--n-arrow-color":O,"--n-arrow-color-disabled":S,"--n-loading-color":D,"--n-color-active-warning":ze,"--n-box-shadow-focus-warning":De,"--n-box-shadow-active-warning":Qe,"--n-box-shadow-hover-warning":Ye,"--n-border-warning":Xe,"--n-border-focus-warning":Ge,"--n-border-hover-warning":et,"--n-border-active-warning":tt,"--n-color-active-error":nt,"--n-box-shadow-focus-error":ot,"--n-box-shadow-active-error":At,"--n-box-shadow-hover-error":jt,"--n-border-error":Et,"--n-border-focus-error":Lt,"--n-border-hover-error":Vt,"--n-border-active-error":Dt,"--n-clear-size":Ht,"--n-clear-color":Nt,"--n-clear-color-hover":Ut,"--n-clear-color-pressed":Wt,"--n-arrow-size":Kt}}),de=Fe?Ze("internal-selection",I(()=>e.size[0]),Me,e):void 0;return{mergedTheme:c,mergedClearable:p,patternInputFocused:v,filterablePlaceholder:z,label:T,selected:R,showTagsPanel:y,isComposing:ae,counterRef:i,counterWrapperRef:b,patternInputMirrorRef:o,patternInputRef:s,selfRef:d,multipleElRef:f,singleElRef:x,patternInputWrapperRef:C,overflowRef:h,inputTagElRef:F,handleMouseDown:J,handleFocusin:se,handleClear:oe,handleMouseEnter:he,handleMouseLeave:ve,handleDeleteOption:K,handlePatternKeyDown:X,handlePatternInputInput:xe,handlePatternInputBlur:Z,handlePatternInputFocus:ge,handleMouseEnterCounter:Te,handleMouseLeaveCounter:Ie,handleFocusout:V,handleCompositionEnd:be,handleCompositionStart:ee,onPopoverUpdateShow:_e,focus:re,focusInput:q,blur:pe,blurInput:j,updateCounter:a,getCounter:m,getTail:ie,renderLabel:e.renderLabel,cssVars:Fe?void 0:Me,themeClass:de==null?void 0:de.themeClass,onRender:de==null?void 0:de.onRender}},render(){const{status:e,multiple:o,size:s,disabled:d,filterable:f,maxTagCount:x,bordered:C,clsPrefix:i,onRender:b,renderTag:h,renderLabel:F}=this;b==null||b();const y=x==="responsive",v=typeof x=="number",g=y||v,c=n(tn,null,{default:()=>n(en,{clsPrefix:i,loading:this.loading,showArrow:this.showArrow,showClear:this.mergedClearable&&this.selected,onClear:this.handleClear},{default:()=>{var z,T;return(T=(z=this.$slots).arrow)===null||T===void 0?void 0:T.call(z)}})});let p;if(o){const{labelField:z}=this,T=V=>n("div",{class:`${i}-base-selection-tag-wrapper`,key:V.value},h?h({option:V,handleClose:()=>this.handleDeleteOption(V)}):n(at,{size:s,closable:!V.disabled,disabled:d,onClose:()=>this.handleDeleteOption(V),internalCloseIsButtonTag:!1,internalCloseFocusable:!1},{default:()=>F?F(V,!0):We(V[z],V,!0)})),R=()=>(v?this.selectedOptions.slice(0,x):this.selectedOptions).map(T),W=f?n("div",{class:`${i}-base-selection-input-tag`,ref:"inputTagElRef",key:"__input-tag__"},n("input",Object.assign({},this.inputProps,{ref:"patternInputRef",tabindex:-1,disabled:d,value:this.pattern,autofocus:this.autofocus,class:`${i}-base-selection-input-tag__input`,onBlur:this.handlePatternInputBlur,onFocus:this.handlePatternInputFocus,onKeydown:this.handlePatternKeyDown,onInput:this.handlePatternInputInput,onCompositionstart:this.handleCompositionStart,onCompositionend:this.handleCompositionEnd})),n("span",{ref:"patternInputMirrorRef",class:`${i}-base-selection-input-tag__mirror`},this.pattern)):null,Y=y?()=>n("div",{class:`${i}-base-selection-tag-wrapper`,ref:"counterWrapperRef"},n(at,{size:s,ref:"counterRef",onMouseenter:this.handleMouseEnterCounter,onMouseleave:this.handleMouseLeaveCounter,disabled:d})):void 0;let H;if(v){const V=this.selectedOptions.length-x;V>0&&(H=n("div",{class:`${i}-base-selection-tag-wrapper`,key:"__counter__"},n(at,{size:s,ref:"counterRef",onMouseenter:this.handleMouseEnterCounter,disabled:d},{default:()=>`+${V}`})))}const _=y?f?n(gt,{ref:"overflowRef",updateCounter:this.updateCounter,getCounter:this.getCounter,getTail:this.getTail,style:{width:"100%",display:"flex",overflow:"hidden"}},{default:R,counter:Y,tail:()=>W}):n(gt,{ref:"overflowRef",updateCounter:this.updateCounter,getCounter:this.getCounter,style:{width:"100%",display:"flex",overflow:"hidden"}},{default:R,counter:Y}):v?R().concat(H):R(),L=g?()=>n("div",{class:`${i}-base-selection-popover`},y?R():this.selectedOptions.map(T)):void 0,N=g?{show:this.showTagsPanel,trigger:"hover",overlap:!0,placement:"top",width:"trigger",onUpdateShow:this.onPopoverUpdateShow,theme:this.mergedTheme.peers.Popover,themeOverrides:this.mergedTheme.peerOverrides.Popover}:null,ne=(this.selected?!1:this.active?!this.pattern&&!this.isComposing:!0)?n("div",{class:`${i}-base-selection-placeholder ${i}-base-selection-overlay`},n("div",{class:`${i}-base-selection-placeholder__inner`},this.placeholder)):null,se=f?n("div",{ref:"patternInputWrapperRef",class:`${i}-base-selection-tags`},_,y?null:W,c):n("div",{ref:"multipleElRef",class:`${i}-base-selection-tags`,tabindex:d?void 0:0},_,c);p=n(lt,null,g?n(Pt,Object.assign({},N,{scrollable:!0,style:"max-height: calc(var(--v-target-height) * 6.6);"}),{trigger:()=>se,default:L}):se,ne)}else if(f){const z=this.pattern||this.isComposing,T=this.active?!z:!this.selected,R=this.active?!1:this.selected;p=n("div",{ref:"patternInputWrapperRef",class:`${i}-base-selection-label`},n("input",Object.assign({},this.inputProps,{ref:"patternInputRef",class:`${i}-base-selection-input`,value:this.active?this.pattern:"",placeholder:"",readonly:d,disabled:d,tabindex:-1,autofocus:this.autofocus,onFocus:this.handlePatternInputFocus,onBlur:this.handlePatternInputBlur,onInput:this.handlePatternInputInput,onCompositionstart:this.handleCompositionStart,onCompositionend:this.handleCompositionEnd})),R?n("div",{class:`${i}-base-selection-label__render-label ${i}-base-selection-overlay`,key:"input"},n("div",{class:`${i}-base-selection-overlay__wrapper`},h?h({option:this.selectedOption,handleClose:()=>{}}):F?F(this.selectedOption,!0):We(this.label,this.selectedOption,!0))):null,T?n("div",{class:`${i}-base-selection-placeholder ${i}-base-selection-overlay`,key:"placeholder"},n("div",{class:`${i}-base-selection-overlay__wrapper`},this.filterablePlaceholder)):null,c)}else p=n("div",{ref:"singleElRef",class:`${i}-base-selection-label`,tabindex:this.disabled?void 0:0},this.label!==void 0?n("div",{class:`${i}-base-selection-input`,title:nn(this.label),key:"input"},n("div",{class:`${i}-base-selection-input__content`},h?h({option:this.selectedOption,handleClose:()=>{}}):F?F(this.selectedOption,!0):We(this.label,this.selectedOption,!0))):n("div",{class:`${i}-base-selection-placeholder ${i}-base-selection-overlay`,key:"placeholder"},n("div",{class:`${i}-base-selection-placeholder__inner`},this.placeholder)),c);return n("div",{ref:"selfRef",class:[`${i}-base-selection`,this.themeClass,e&&`${i}-base-selection--${e}-status`,{[`${i}-base-selection--active`]:this.active,[`${i}-base-selection--selected`]:this.selected||this.active&&this.pattern,[`${i}-base-selection--disabled`]:this.disabled,[`${i}-base-selection--multiple`]:this.multiple,[`${i}-base-selection--focus`]:this.focused}],style:this.cssVars,onClick:this.onClick,onMouseenter:this.handleMouseEnter,onMouseleave:this.handleMouseLeave,onKeydown:this.onKeydown,onFocusin:this.handleFocusin,onFocusout:this.handleFocusout,onMousedown:this.handleMouseDown},p,C?n("div",{class:`${i}-base-selection__border`}):null,C?n("div",{class:`${i}-base-selection__state-border`}):null)}});function Je(e){return e.type==="group"}function It(e){return e.type==="ignored"}function rt(e,o){try{return!!(1+o.toString().toLowerCase().indexOf(e.trim().toLowerCase()))}catch{return!1}}function _t(e,o){return{getIsGroup:Je,getIgnored:It,getKey(d){return Je(d)?d.name||d.key||"key-required":d[e]},getChildren(d){return d[o]}}}function zn(e,o,s,d){if(!o)return e;function f(x){if(!Array.isArray(x))return[];const C=[];for(const i of x)if(Je(i)){const b=f(i[d]);b.length&&C.push(Object.assign({},i,{[d]:b}))}else{if(It(i))continue;o(s,i)&&C.push(i)}return C}return f(e)}function Rn(e,o,s){const d=new Map;return e.forEach(f=>{Je(f)?f[s].forEach(x=>{d.set(x[o],x)}):d.set(f[o],f)}),d}const $t=on("n-popselect"),Tn=k("popselect-menu",`
 box-shadow: var(--n-menu-box-shadow);
`),ut={multiple:Boolean,value:{type:[String,Number,Array],default:null},cancelable:Boolean,options:{type:Array,default:()=>[]},size:{type:String,default:"medium"},scrollable:Boolean,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onMouseenter:Function,onMouseleave:Function,renderLabel:Function,showCheckmark:{type:Boolean,default:void 0},nodeProps:Function,virtualScroll:Boolean,onChange:[Function,Array]},yt=rn(ut),In=le({name:"PopselectPanel",props:ut,setup(e){const o=an($t),{mergedClsPrefixRef:s,inlineThemeDisabled:d}=dt(e),f=Ce("Popselect","-pop-select",Tn,Ot,o.props,s),x=I(()=>Bt(e.options,_t("value","children")));function C(v,g){const{onUpdateValue:c,"onUpdate:value":p,onChange:z}=e;c&&E(c,v,g),p&&E(p,v,g),z&&E(z,v,g)}function i(v){h(v.key)}function b(v){st(v,"action")||v.preventDefault()}function h(v){const{value:{getNode:g}}=x;if(e.multiple)if(Array.isArray(e.value)){const c=[],p=[];let z=!0;e.value.forEach(T=>{if(T===v){z=!1;return}const R=g(T);R&&(c.push(R.key),p.push(R.rawNode))}),z&&(c.push(v),p.push(g(v).rawNode)),C(c,p)}else{const c=g(v);c&&C([v],[c.rawNode])}else if(e.value===v&&e.cancelable)C(null,null);else{const c=g(v);c&&C(v,c.rawNode);const{"onUpdate:show":p,onUpdateShow:z}=o.props;p&&E(p,!1),z&&E(z,!1),o.setShow(!1)}Ue(()=>{o.syncPosition()})}Ne(we(e,"options"),()=>{Ue(()=>{o.syncPosition()})});const F=I(()=>{const{self:{menuBoxShadow:v}}=f.value;return{"--n-menu-box-shadow":v}}),y=d?Ze("select",void 0,F,o.props):void 0;return{mergedTheme:o.mergedThemeRef,mergedClsPrefix:s,treeMate:x,handleToggle:i,handleMenuMousedown:b,cssVars:d?void 0:F,themeClass:y==null?void 0:y.themeClass,onRender:y==null?void 0:y.onRender}},render(){var e;return(e=this.onRender)===null||e===void 0||e.call(this),n(Tt,{clsPrefix:this.mergedClsPrefix,focusable:!0,nodeProps:this.nodeProps,class:[`${this.mergedClsPrefix}-popselect-menu`,this.themeClass],style:this.cssVars,theme:this.mergedTheme.peers.InternalSelectMenu,themeOverrides:this.mergedTheme.peerOverrides.InternalSelectMenu,multiple:this.multiple,treeMate:this.treeMate,size:this.size,value:this.value,virtualScroll:this.virtualScroll,scrollable:this.scrollable,renderLabel:this.renderLabel,onToggle:this.handleToggle,onMouseenter:this.onMouseenter,onMouseleave:this.onMouseenter,onMousedown:this.handleMenuMousedown,showCheckmark:this.showCheckmark},{action:()=>{var o,s;return((s=(o=this.$slots).action)===null||s===void 0?void 0:s.call(o))||[]},empty:()=>{var o,s;return((s=(o=this.$slots).empty)===null||s===void 0?void 0:s.call(o))||[]}})}}),_n=Object.assign(Object.assign(Object.assign(Object.assign({},Ce.props),zt(ct,["showArrow","arrow"])),{placement:Object.assign(Object.assign({},ct.placement),{default:"bottom"}),trigger:{type:String,default:"hover"}}),ut),$n=le({name:"Popselect",props:_n,inheritAttrs:!1,__popover__:!0,setup(e){const o=Ce("Popselect","-popselect",void 0,Ot,e),s=P(null);function d(){var C;(C=s.value)===null||C===void 0||C.syncPosition()}function f(C){var i;(i=s.value)===null||i===void 0||i.setShow(C)}return dn($t,{props:e,mergedThemeRef:o,syncPosition:d,setShow:f}),Object.assign(Object.assign({},{syncPosition:d,setShow:f}),{popoverInstRef:s,mergedTheme:o})},render(){const{mergedTheme:e}=this,o={theme:e.peers.Popover,themeOverrides:e.peerOverrides.Popover,builtinThemeOverrides:{padding:"0"},ref:"popoverInstRef",internalRenderBody:(s,d,f,x,C)=>{const{$attrs:i}=this;return n(In,Object.assign({},i,{class:[i.class,s],style:[i.style,f]},ln(this.$props,yt),{ref:sn(d),onMouseenter:vt([x,i.onMouseenter]),onMouseleave:vt([C,i.onMouseleave])}),{action:()=>{var b,h;return(h=(b=this.$slots).action)===null||h===void 0?void 0:h.call(b)},empty:()=>{var b,h;return(h=(b=this.$slots).empty)===null||h===void 0?void 0:h.call(b)}})}};return n(Pt,Object.assign({},zt(this.$props,yt),o,{internalDeactivateImmediately:!0}),{trigger:()=>{var s,d;return(d=(s=this.$slots).default)===null||d===void 0?void 0:d.call(s)}})}}),An=ce([k("select",`
 z-index: auto;
 outline: none;
 width: 100%;
 position: relative;
 `),k("select-menu",`
 margin: 4px 0;
 box-shadow: var(--n-menu-box-shadow);
 `,[un({originalTransition:"background-color .3s var(--n-bezier), box-shadow .3s var(--n-bezier)"})])]),jn=Object.assign(Object.assign({},Ce.props),{to:qe.propTo,bordered:{type:Boolean,default:void 0},clearable:Boolean,clearFilterAfterSelect:{type:Boolean,default:!0},options:{type:Array,default:()=>[]},defaultValue:{type:[String,Number,Array],default:null},value:[String,Number,Array],placeholder:String,menuProps:Object,multiple:Boolean,size:String,filterable:Boolean,disabled:{type:Boolean,default:void 0},remote:Boolean,loading:Boolean,filter:Function,placement:{type:String,default:"bottom-start"},widthMode:{type:String,default:"trigger"},tag:Boolean,onCreate:Function,fallbackOption:{type:[Function,Boolean],default:void 0},show:{type:Boolean,default:void 0},showArrow:{type:Boolean,default:!0},maxTagCount:[Number,String],consistentMenuWidth:{type:Boolean,default:!0},virtualScroll:{type:Boolean,default:!0},labelField:{type:String,default:"label"},valueField:{type:String,default:"value"},childrenField:{type:String,default:"children"},renderLabel:Function,renderOption:Function,renderTag:Function,"onUpdate:value":[Function,Array],inputProps:Object,nodeProps:Function,ignoreComposition:{type:Boolean,default:!0},showOnFocus:Boolean,onUpdateValue:[Function,Array],onBlur:[Function,Array],onClear:[Function,Array],onFocus:[Function,Array],onScroll:[Function,Array],onSearch:[Function,Array],onUpdateShow:[Function,Array],"onUpdate:show":[Function,Array],displayDirective:{type:String,default:"show"},resetMenuOnOptionsChange:{type:Boolean,default:!0},status:String,showCheckmark:{type:Boolean,default:!0},onChange:[Function,Array],items:Array}),En=le({name:"Select",props:jn,setup(e){const{mergedClsPrefixRef:o,mergedBorderedRef:s,namespaceRef:d,inlineThemeDisabled:f}=dt(e),x=Ce("Select","-select",An,cn,e,o),C=P(e.defaultValue),i=we(e,"value"),b=Ke(i,C),h=P(!1),F=P(""),y=I(()=>{const{valueField:t,childrenField:l}=e,w=_t(t,l);return Bt(V.value,w)}),v=I(()=>Rn(ne.value,e.valueField,e.childrenField)),g=P(!1),c=Ke(we(e,"show"),g),p=P(null),z=P(null),T=P(null),{localeRef:R}=Rt("Select"),W=I(()=>{var t;return(t=e.placeholder)!==null&&t!==void 0?t:R.value.placeholder}),Y=fn(e,["items","options"]),H=[],_=P([]),L=P([]),N=P(new Map),fe=I(()=>{const{fallbackOption:t}=e;if(t===void 0){const{labelField:l,valueField:w}=e;return B=>({[l]:String(B),[w]:B})}return t===!1?!1:l=>Object.assign(t(l),{value:l})}),ne=I(()=>L.value.concat(_.value).concat(Y.value)),se=I(()=>{const{filter:t}=e;if(t)return t;const{labelField:l,valueField:w}=e;return(B,M)=>{if(!M)return!1;const O=M[l];if(typeof O=="string")return rt(B,O);const S=M[w];return typeof S=="string"?rt(B,S):typeof S=="number"?rt(B,String(S)):!1}}),V=I(()=>{if(e.remote)return Y.value;{const{value:t}=ne,{value:l}=F;return!l.length||!e.filterable?t:zn(t,se.value,l,e.childrenField)}});function oe(t){const l=e.remote,{value:w}=N,{value:B}=v,{value:M}=fe,O=[];return t.forEach(S=>{if(B.has(S))O.push(B.get(S));else if(l&&w.has(S))O.push(w.get(S));else if(M){const D=M(S);D&&O.push(D)}}),O}const he=I(()=>{if(e.multiple){const{value:t}=b;return Array.isArray(t)?oe(t):[]}return null}),ve=I(()=>{const{value:t}=b;return!e.multiple&&!Array.isArray(t)?t===null?null:oe([t])[0]||null:null}),J=hn(e),{mergedSizeRef:K,mergedDisabledRef:X,mergedStatusRef:ae}=J;function A(t,l){const{onChange:w,"onUpdate:value":B,onUpdateValue:M}=e,{nTriggerFormChange:O,nTriggerFormInput:S}=J;w&&E(w,t,l),M&&E(M,t,l),B&&E(B,t,l),C.value=t,O(),S()}function xe(t){const{onBlur:l}=e,{nTriggerFormBlur:w}=J;l&&E(l,t),w()}function ee(){const{onClear:t}=e;t&&E(t)}function be(t){const{onFocus:l,showOnFocus:w}=e,{nTriggerFormFocus:B}=J;l&&E(l,t),B(),w&&q()}function ge(t){const{onSearch:l}=e;l&&E(l,t)}function Z(t){const{onScroll:l}=e;l&&E(l,t)}function pe(){var t;const{remote:l,multiple:w}=e;if(l){const{value:B}=N;if(w){const{valueField:M}=e;(t=he.value)===null||t===void 0||t.forEach(O=>{B.set(O[M],O)})}else{const M=ve.value;M&&B.set(M[e.valueField],M)}}}function re(t){const{onUpdateShow:l,"onUpdate:show":w}=e;l&&E(l,t),w&&E(w,t),g.value=t}function q(){X.value||(re(!0),g.value=!0,e.filterable&&Oe())}function j(){re(!1)}function a(){F.value="",L.value=H}const m=P(!1);function ie(){e.filterable&&(m.value=!0)}function te(){e.filterable&&(m.value=!1,c.value||a())}function me(){X.value||(c.value?e.filterable?Oe():j():q())}function Te(t){var l,w;!((w=(l=T.value)===null||l===void 0?void 0:l.selfRef)===null||w===void 0)&&w.contains(t.relatedTarget)||(h.value=!1,xe(t),j())}function Ie(t){be(t),h.value=!0}function _e(t){h.value=!0}function Fe(t){var l;!((l=p.value)===null||l===void 0)&&l.$el.contains(t.relatedTarget)||(h.value=!1,xe(t),j())}function Me(){var t;(t=p.value)===null||t===void 0||t.focus(),j()}function de(t){var l;c.value&&(!((l=p.value)===null||l===void 0)&&l.$el.contains(xn(t))||j())}function r(t){if(!Array.isArray(t))return[];if(fe.value)return Array.from(t);{const{remote:l}=e,{value:w}=v;if(l){const{value:B}=N;return t.filter(M=>w.has(M)||B.has(M))}else return t.filter(B=>w.has(B))}}function u(t){$(t.rawNode)}function $(t){if(X.value)return;const{tag:l,remote:w,clearFilterAfterSelect:B,valueField:M}=e;if(l&&!w){const{value:O}=L,S=O[0]||null;if(S){const D=_.value;D.length?D.push(S):_.value=[S],L.value=H}}if(w&&N.value.set(t[M],t),e.multiple){const O=r(b.value),S=O.findIndex(D=>D===t[M]);if(~S){if(O.splice(S,1),l&&!w){const D=Se(t[M]);~D&&(_.value.splice(D,1),B&&(F.value=""))}}else O.push(t[M]),B&&(F.value="");A(O,oe(O))}else{if(l&&!w){const O=Se(t[M]);~O?_.value=[_.value[O]]:_.value=H}Pe(),j(),A(t[M],t)}}function Se(t){return _.value.findIndex(w=>w[e.valueField]===t)}function $e(t){c.value||q();const{value:l}=t.target;F.value=l;const{tag:w,remote:B}=e;if(ge(l),w&&!B){if(!l){L.value=H;return}const{onCreate:M}=e,O=M?M(l):{[e.labelField]:l,[e.valueField]:l},{valueField:S}=e;Y.value.some(D=>D[S]===O[S])||_.value.some(D=>D[S]===O[S])?L.value=H:L.value=[O]}}function Ae(t){t.stopPropagation();const{multiple:l}=e;!l&&e.filterable&&j(),ee(),l?A([],[]):A(null,null)}function je(t){!st(t,"action")&&!st(t,"empty")&&t.preventDefault()}function Ee(t){Z(t)}function ke(t){var l,w,B,M,O;switch(t.key){case" ":if(e.filterable)break;t.preventDefault();case"Enter":if(!(!((l=p.value)===null||l===void 0)&&l.isComposing)){if(c.value){const S=(w=T.value)===null||w===void 0?void 0:w.getPendingTmNode();S?u(S):e.filterable||(j(),Pe())}else if(q(),e.tag&&m.value){const S=L.value[0];if(S){const D=S[e.valueField],{value:ze}=b;e.multiple&&Array.isArray(ze)&&ze.some(De=>De===D)||$(S)}}}t.preventDefault();break;case"ArrowUp":if(t.preventDefault(),e.loading)return;c.value&&((B=T.value)===null||B===void 0||B.prev());break;case"ArrowDown":if(t.preventDefault(),e.loading)return;c.value?(M=T.value)===null||M===void 0||M.next():q();break;case"Escape":c.value&&(yn(t),j()),(O=p.value)===null||O===void 0||O.focus();break}}function Pe(){var t;(t=p.value)===null||t===void 0||t.focus()}function Oe(){var t;(t=p.value)===null||t===void 0||t.focusInput()}function Le(){var t;c.value&&((t=z.value)===null||t===void 0||t.syncPosition())}pe(),Ne(we(e,"options"),pe);const Ve={focus:()=>{var t;(t=p.value)===null||t===void 0||t.focus()},blur:()=>{var t;(t=p.value)===null||t===void 0||t.blur()}},Be=I(()=>{const{self:{menuBoxShadow:t}}=x.value;return{"--n-menu-box-shadow":t}}),ue=f?Ze("select",void 0,Be,e):void 0;return Object.assign(Object.assign({},Ve),{mergedStatus:ae,mergedClsPrefix:o,mergedBordered:s,namespace:d,treeMate:y,isMounted:vn(),triggerRef:p,menuRef:T,pattern:F,uncontrolledShow:g,mergedShow:c,adjustedTo:qe(e),uncontrolledValue:C,mergedValue:b,followerRef:z,localizedPlaceholder:W,selectedOption:ve,selectedOptions:he,mergedSize:K,mergedDisabled:X,focused:h,activeWithoutMenuOpen:m,inlineThemeDisabled:f,onTriggerInputFocus:ie,onTriggerInputBlur:te,handleTriggerOrMenuResize:Le,handleMenuFocus:_e,handleMenuBlur:Fe,handleMenuTabOut:Me,handleTriggerClick:me,handleToggle:u,handleDeleteOption:$,handlePatternInput:$e,handleClear:Ae,handleTriggerBlur:Te,handleTriggerFocus:Ie,handleKeydown:ke,handleMenuAfterLeave:a,handleMenuClickOutside:de,handleMenuScroll:Ee,handleMenuKeydown:ke,handleMenuMousedown:je,mergedTheme:x,cssVars:f?void 0:Be,themeClass:ue==null?void 0:ue.themeClass,onRender:ue==null?void 0:ue.onRender})},render(){return n("div",{class:`${this.mergedClsPrefix}-select`},n(bn,null,{default:()=>[n(gn,null,{default:()=>n(Bn,{ref:"triggerRef",inlineThemeDisabled:this.inlineThemeDisabled,status:this.mergedStatus,inputProps:this.inputProps,clsPrefix:this.mergedClsPrefix,showArrow:this.showArrow,maxTagCount:this.maxTagCount,bordered:this.mergedBordered,active:this.activeWithoutMenuOpen||this.mergedShow,pattern:this.pattern,placeholder:this.localizedPlaceholder,selectedOption:this.selectedOption,selectedOptions:this.selectedOptions,multiple:this.multiple,renderTag:this.renderTag,renderLabel:this.renderLabel,filterable:this.filterable,clearable:this.clearable,disabled:this.mergedDisabled,size:this.mergedSize,theme:this.mergedTheme.peers.InternalSelection,labelField:this.labelField,valueField:this.valueField,themeOverrides:this.mergedTheme.peerOverrides.InternalSelection,loading:this.loading,focused:this.focused,onClick:this.handleTriggerClick,onDeleteOption:this.handleDeleteOption,onPatternInput:this.handlePatternInput,onClear:this.handleClear,onBlur:this.handleTriggerBlur,onFocus:this.handleTriggerFocus,onKeydown:this.handleKeydown,onPatternBlur:this.onTriggerInputBlur,onPatternFocus:this.onTriggerInputFocus,onResize:this.handleTriggerOrMenuResize,ignoreComposition:this.ignoreComposition},{arrow:()=>{var e,o;return[(o=(e=this.$slots).arrow)===null||o===void 0?void 0:o.call(e)]}})}),n(pn,{ref:"followerRef",show:this.mergedShow,to:this.adjustedTo,teleportDisabled:this.adjustedTo===qe.tdkey,containerClass:this.namespace,width:this.consistentMenuWidth?"target":void 0,minWidth:"target",placement:this.placement},{default:()=>n(mn,{name:"fade-in-scale-up-transition",appear:this.isMounted,onAfterLeave:this.handleMenuAfterLeave},{default:()=>{var e,o,s;return this.mergedShow||this.displayDirective==="show"?((e=this.onRender)===null||e===void 0||e.call(this),wn(n(Tt,Object.assign({},this.menuProps,{ref:"menuRef",onResize:this.handleTriggerOrMenuResize,inlineThemeDisabled:this.inlineThemeDisabled,virtualScroll:this.consistentMenuWidth&&this.virtualScroll,class:[`${this.mergedClsPrefix}-select-menu`,this.themeClass,(o=this.menuProps)===null||o===void 0?void 0:o.class],clsPrefix:this.mergedClsPrefix,focusable:!0,labelField:this.labelField,valueField:this.valueField,autoPending:!0,nodeProps:this.nodeProps,theme:this.mergedTheme.peers.InternalSelectMenu,themeOverrides:this.mergedTheme.peerOverrides.InternalSelectMenu,treeMate:this.treeMate,multiple:this.multiple,size:"medium",renderOption:this.renderOption,renderLabel:this.renderLabel,value:this.mergedValue,style:[(s=this.menuProps)===null||s===void 0?void 0:s.style,this.cssVars],onToggle:this.handleToggle,onScroll:this.handleMenuScroll,onFocus:this.handleMenuFocus,onBlur:this.handleMenuBlur,onKeydown:this.handleMenuKeydown,onTabOut:this.handleMenuTabOut,onMousedown:this.handleMenuMousedown,show:this.mergedShow,showCheckmark:this.showCheckmark,resetMenuOnOptionsChange:this.resetMenuOnOptionsChange}),{empty:()=>{var d,f;return[(f=(d=this.$slots).empty)===null||f===void 0?void 0:f.call(d)]},action:()=>{var d,f;return[(f=(d=this.$slots).action)===null||f===void 0?void 0:f.call(d)]}}),this.displayDirective==="show"?[[Cn,this.mergedShow],[ft,this.handleMenuClickOutside,void 0,{capture:!0}]]:[[ft,this.handleMenuClickOutside,void 0,{capture:!0}]])):null}})})]}))}});function Ln(e,o,s){let d=!1,f=!1,x=1,C=o;if(o===1)return{hasFastBackward:!1,hasFastForward:!1,fastForwardTo:C,fastBackwardTo:x,items:[{type:"page",label:1,active:e===1,mayBeFastBackward:!1,mayBeFastForward:!1}]};if(o===2)return{hasFastBackward:!1,hasFastForward:!1,fastForwardTo:C,fastBackwardTo:x,items:[{type:"page",label:1,active:e===1,mayBeFastBackward:!1,mayBeFastForward:!1},{type:"page",label:2,active:e===2,mayBeFastBackward:!0,mayBeFastForward:!1}]};const i=1,b=o;let h=e,F=e;const y=(s-5)/2;F+=Math.ceil(y),F=Math.min(Math.max(F,i+s-3),b-2),h-=Math.floor(y),h=Math.max(Math.min(h,b-s+3),i+2);let v=!1,g=!1;h>i+2&&(v=!0),F<b-2&&(g=!0);const c=[];c.push({type:"page",label:1,active:e===1,mayBeFastBackward:!1,mayBeFastForward:!1}),v?(d=!0,x=h-1,c.push({type:"fast-backward",active:!1,label:void 0,options:Ft(i+1,h-1)})):b>=i+1&&c.push({type:"page",label:i+1,mayBeFastBackward:!0,mayBeFastForward:!1,active:e===i+1});for(let p=h;p<=F;++p)c.push({type:"page",label:p,mayBeFastBackward:!1,mayBeFastForward:!1,active:e===p});return g?(f=!0,C=F+1,c.push({type:"fast-forward",active:!1,label:void 0,options:Ft(F+1,b-1)})):F===b-2&&c[c.length-1].label!==b-1&&c.push({type:"page",mayBeFastForward:!0,mayBeFastBackward:!1,label:b-1,active:e===b-1}),c[c.length-1].label!==b&&c.push({type:"page",mayBeFastForward:!1,mayBeFastBackward:!1,label:b,active:e===b}),{hasFastBackward:d,hasFastForward:f,fastBackwardTo:x,fastForwardTo:C,items:c}}function Ft(e,o){const s=[];for(let d=e;d<=o;++d)s.push({label:`${d}`,value:d});return s}const Mt=`
 background: var(--n-item-color-hover);
 color: var(--n-item-text-color-hover);
 border: var(--n-item-border-hover);
`,St=[G("button",`
 background: var(--n-button-color-hover);
 border: var(--n-button-border-hover);
 color: var(--n-button-icon-color-hover);
 `)],Vn=k("pagination",`
 display: flex;
 vertical-align: middle;
 font-size: var(--n-item-font-size);
 flex-wrap: nowrap;
`,[k("pagination-prefix",`
 display: flex;
 align-items: center;
 margin: var(--n-prefix-margin);
 `),k("pagination-suffix",`
 display: flex;
 align-items: center;
 margin: var(--n-suffix-margin);
 `),ce("> *:not(:first-child)",`
 margin: var(--n-item-margin);
 `),k("select",`
 width: var(--n-select-width);
 `),ce("&.transition-disabled",[k("pagination-item","transition: none!important;")]),k("pagination-quick-jumper",`
 white-space: nowrap;
 display: flex;
 color: var(--n-jumper-text-color);
 transition: color .3s var(--n-bezier);
 align-items: center;
 font-size: var(--n-jumper-font-size);
 `,[k("input",`
 margin: var(--n-input-margin);
 width: var(--n-input-width);
 `)]),k("pagination-item",`
 position: relative;
 cursor: pointer;
 user-select: none;
 -webkit-user-select: none;
 display: flex;
 align-items: center;
 justify-content: center;
 box-sizing: border-box;
 min-width: var(--n-item-size);
 height: var(--n-item-size);
 padding: var(--n-item-padding);
 background-color: var(--n-item-color);
 color: var(--n-item-text-color);
 border-radius: var(--n-item-border-radius);
 border: var(--n-item-border);
 fill: var(--n-button-icon-color);
 transition:
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 fill .3s var(--n-bezier);
 `,[G("button",`
 background: var(--n-button-color);
 color: var(--n-button-icon-color);
 border: var(--n-button-border);
 padding: 0;
 `,[k("base-icon",`
 font-size: var(--n-button-icon-size);
 `)]),it("disabled",[G("hover",Mt,St),ce("&:hover",Mt,St),ce("&:active",`
 background: var(--n-item-color-pressed);
 color: var(--n-item-text-color-pressed);
 border: var(--n-item-border-pressed);
 `,[G("button",`
 background: var(--n-button-color-pressed);
 border: var(--n-button-border-pressed);
 color: var(--n-button-icon-color-pressed);
 `)]),G("active",`
 background: var(--n-item-color-active);
 color: var(--n-item-text-color-active);
 border: var(--n-item-border-active);
 `,[ce("&:hover",`
 background: var(--n-item-color-active-hover);
 `)])]),G("disabled",`
 cursor: not-allowed;
 color: var(--n-item-text-color-disabled);
 `,[G("active, button",`
 background-color: var(--n-item-color-disabled);
 border: var(--n-item-border-disabled);
 `)])]),G("disabled",`
 cursor: not-allowed;
 `,[k("pagination-quick-jumper",`
 color: var(--n-jumper-text-color-disabled);
 `)]),G("simple",`
 display: flex;
 align-items: center;
 flex-wrap: nowrap;
 `,[k("pagination-quick-jumper",[k("input",`
 margin: 0;
 `)])])]),Dn=Object.assign(Object.assign({},Ce.props),{simple:Boolean,page:Number,defaultPage:{type:Number,default:1},itemCount:Number,pageCount:Number,defaultPageCount:{type:Number,default:1},showSizePicker:Boolean,pageSize:Number,defaultPageSize:Number,pageSizes:{type:Array,default(){return[10]}},showQuickJumper:Boolean,size:{type:String,default:"medium"},disabled:Boolean,pageSlot:{type:Number,default:9},selectProps:Object,prev:Function,next:Function,goto:Function,prefix:Function,suffix:Function,label:Function,displayOrder:{type:Array,default:["pages","size-picker","quick-jumper"]},to:qe.propTo,"onUpdate:page":[Function,Array],onUpdatePage:[Function,Array],"onUpdate:pageSize":[Function,Array],onUpdatePageSize:[Function,Array],onPageSizeChange:[Function,Array],onChange:[Function,Array]}),Wn=le({name:"Pagination",props:Dn,setup(e){const{mergedComponentPropsRef:o,mergedClsPrefixRef:s,inlineThemeDisabled:d,mergedRtlRef:f}=dt(e),x=Ce("Pagination","-pagination",Vn,Fn,e,s),{localeRef:C}=Rt("Pagination"),i=P(null),b=P(e.defaultPage),F=P((()=>{const{defaultPageSize:a}=e;if(a!==void 0)return a;const m=e.pageSizes[0];return typeof m=="number"?m:m.value||10})()),y=Ke(we(e,"page"),b),v=Ke(we(e,"pageSize"),F),g=I(()=>{const{itemCount:a}=e;if(a!==void 0)return Math.max(1,Math.ceil(a/v.value));const{pageCount:m}=e;return m!==void 0?Math.max(m,1):1}),c=P("");He(()=>{e.simple,c.value=String(y.value)});const p=P(!1),z=P(!1),T=P(!1),R=P(!1),W=()=>{e.disabled||(p.value=!0,J())},Y=()=>{e.disabled||(p.value=!1,J())},H=()=>{z.value=!0,J()},_=()=>{z.value=!1,J()},L=a=>{K(a)},N=I(()=>Ln(y.value,g.value,e.pageSlot));He(()=>{N.value.hasFastBackward?N.value.hasFastForward||(p.value=!1,T.value=!1):(z.value=!1,R.value=!1)});const fe=I(()=>{const a=C.value.selectionSuffix;return e.pageSizes.map(m=>typeof m=="number"?{label:`${m} / ${a}`,value:m}:m)}),ne=I(()=>{var a,m;return((m=(a=o==null?void 0:o.value)===null||a===void 0?void 0:a.Pagination)===null||m===void 0?void 0:m.inputSize)||bt(e.size)}),se=I(()=>{var a,m;return((m=(a=o==null?void 0:o.value)===null||a===void 0?void 0:a.Pagination)===null||m===void 0?void 0:m.selectSize)||bt(e.size)}),V=I(()=>(y.value-1)*v.value),oe=I(()=>{const a=y.value*v.value-1,{itemCount:m}=e;return m!==void 0&&a>m-1?m-1:a}),he=I(()=>{const{itemCount:a}=e;return a!==void 0?a:(e.pageCount||1)*v.value}),ve=Mn("Pagination",f,s),J=()=>{Ue(()=>{var a;const{value:m}=i;m&&(m.classList.add("transition-disabled"),(a=i.value)===null||a===void 0||a.offsetWidth,m.classList.remove("transition-disabled"))})};function K(a){if(a===y.value)return;const{"onUpdate:page":m,onUpdatePage:ie,onChange:te,simple:me}=e;m&&E(m,a),ie&&E(ie,a),te&&E(te,a),b.value=a,me&&(c.value=String(a))}function X(a){if(a===v.value)return;const{"onUpdate:pageSize":m,onUpdatePageSize:ie,onPageSizeChange:te}=e;m&&E(m,a),ie&&E(ie,a),te&&E(te,a),F.value=a,g.value<y.value&&K(g.value)}function ae(){if(e.disabled)return;const a=Math.min(y.value+1,g.value);K(a)}function A(){if(e.disabled)return;const a=Math.max(y.value-1,1);K(a)}function xe(){if(e.disabled)return;const a=Math.min(N.value.fastForwardTo,g.value);K(a)}function ee(){if(e.disabled)return;const a=Math.max(N.value.fastBackwardTo,1);K(a)}function be(a){X(a)}function ge(){const a=parseInt(c.value);Number.isNaN(a)||(K(Math.max(1,Math.min(a,g.value))),e.simple||(c.value=""))}function Z(){ge()}function pe(a){if(!e.disabled)switch(a.type){case"page":K(a.label);break;case"fast-backward":ee();break;case"fast-forward":xe();break}}function re(a){c.value=a.replace(/\D+/g,"")}He(()=>{y.value,v.value,J()});const q=I(()=>{const{size:a}=e,{self:{buttonBorder:m,buttonBorderHover:ie,buttonBorderPressed:te,buttonIconColor:me,buttonIconColorHover:Te,buttonIconColorPressed:Ie,itemTextColor:_e,itemTextColorHover:Fe,itemTextColorPressed:Me,itemTextColorActive:de,itemTextColorDisabled:r,itemColor:u,itemColorHover:$,itemColorPressed:Se,itemColorActive:$e,itemColorActiveHover:Ae,itemColorDisabled:je,itemBorder:Ee,itemBorderHover:ke,itemBorderPressed:Pe,itemBorderActive:Oe,itemBorderDisabled:Le,itemBorderRadius:Ve,jumperTextColor:Be,jumperTextColorDisabled:ue,buttonColor:t,buttonColorHover:l,buttonColorPressed:w,[Q("itemPadding",a)]:B,[Q("itemMargin",a)]:M,[Q("inputWidth",a)]:O,[Q("selectWidth",a)]:S,[Q("inputMargin",a)]:D,[Q("selectMargin",a)]:ze,[Q("jumperFontSize",a)]:De,[Q("prefixMargin",a)]:Qe,[Q("suffixMargin",a)]:Ye,[Q("itemSize",a)]:Xe,[Q("buttonIconSize",a)]:Ge,[Q("itemFontSize",a)]:et,[`${Q("itemMargin",a)}Rtl`]:tt,[`${Q("inputMargin",a)}Rtl`]:nt},common:{cubicBezierEaseInOut:ot}}=x.value;return{"--n-prefix-margin":Qe,"--n-suffix-margin":Ye,"--n-item-font-size":et,"--n-select-width":S,"--n-select-margin":ze,"--n-input-width":O,"--n-input-margin":D,"--n-input-margin-rtl":nt,"--n-item-size":Xe,"--n-item-text-color":_e,"--n-item-text-color-disabled":r,"--n-item-text-color-hover":Fe,"--n-item-text-color-active":de,"--n-item-text-color-pressed":Me,"--n-item-color":u,"--n-item-color-hover":$,"--n-item-color-disabled":je,"--n-item-color-active":$e,"--n-item-color-active-hover":Ae,"--n-item-color-pressed":Se,"--n-item-border":Ee,"--n-item-border-hover":ke,"--n-item-border-disabled":Le,"--n-item-border-active":Oe,"--n-item-border-pressed":Pe,"--n-item-padding":B,"--n-item-border-radius":Ve,"--n-bezier":ot,"--n-jumper-font-size":De,"--n-jumper-text-color":Be,"--n-jumper-text-color-disabled":ue,"--n-item-margin":M,"--n-item-margin-rtl":tt,"--n-button-icon-size":Ge,"--n-button-icon-color":me,"--n-button-icon-color-hover":Te,"--n-button-icon-color-pressed":Ie,"--n-button-color-hover":l,"--n-button-color":t,"--n-button-color-pressed":w,"--n-button-border":m,"--n-button-border-hover":ie,"--n-button-border-pressed":te}}),j=d?Ze("pagination",I(()=>{let a="";const{size:m}=e;return a+=m[0],a}),q,e):void 0;return{rtlEnabled:ve,mergedClsPrefix:s,locale:C,selfRef:i,mergedPage:y,pageItems:I(()=>N.value.items),mergedItemCount:he,jumperValue:c,pageSizeOptions:fe,mergedPageSize:v,inputSize:ne,selectSize:se,mergedTheme:x,mergedPageCount:g,startIndex:V,endIndex:oe,showFastForwardMenu:T,showFastBackwardMenu:R,fastForwardActive:p,fastBackwardActive:z,handleMenuSelect:L,handleFastForwardMouseenter:W,handleFastForwardMouseleave:Y,handleFastBackwardMouseenter:H,handleFastBackwardMouseleave:_,handleJumperInput:re,handleBackwardClick:A,handleForwardClick:ae,handlePageItemClick:pe,handleSizePickerChange:be,handleQuickJumperChange:Z,cssVars:d?void 0:q,themeClass:j==null?void 0:j.themeClass,onRender:j==null?void 0:j.onRender}},render(){const{$slots:e,mergedClsPrefix:o,disabled:s,cssVars:d,mergedPage:f,mergedPageCount:x,pageItems:C,showSizePicker:i,showQuickJumper:b,mergedTheme:h,locale:F,inputSize:y,selectSize:v,mergedPageSize:g,pageSizeOptions:c,jumperValue:p,simple:z,prev:T,next:R,prefix:W,suffix:Y,label:H,goto:_,handleJumperInput:L,handleSizePickerChange:N,handleBackwardClick:fe,handlePageItemClick:ne,handleForwardClick:se,handleQuickJumperChange:V,onRender:oe}=this;oe==null||oe();const he=e.prefix||W,ve=e.suffix||Y,J=T||e.prev,K=R||e.next,X=H||e.label;return n("div",{ref:"selfRef",class:[`${o}-pagination`,this.themeClass,this.rtlEnabled&&`${o}-pagination--rtl`,s&&`${o}-pagination--disabled`,z&&`${o}-pagination--simple`],style:d},he?n("div",{class:`${o}-pagination-prefix`},he({page:f,pageSize:g,pageCount:x,startIndex:this.startIndex,endIndex:this.endIndex,itemCount:this.mergedItemCount})):null,this.displayOrder.map(ae=>{switch(ae){case"pages":return n(lt,null,n("div",{class:[`${o}-pagination-item`,!J&&`${o}-pagination-item--button`,(f<=1||f>x||s)&&`${o}-pagination-item--disabled`],onClick:fe},J?J({page:f,pageSize:g,pageCount:x,startIndex:this.startIndex,endIndex:this.endIndex,itemCount:this.mergedItemCount}):n(Re,{clsPrefix:o},{default:()=>this.rtlEnabled?n(Ct,null):n(pt,null)})),z?n(lt,null,n("div",{class:`${o}-pagination-quick-jumper`},n(ht,{value:p,onUpdateValue:L,size:y,placeholder:"",disabled:s,theme:h.peers.Input,themeOverrides:h.peerOverrides.Input,onChange:V}))," / ",x):C.map((A,xe)=>{let ee,be,ge;const{type:Z}=A;switch(Z){case"page":const re=A.label;X?ee=X({type:"page",node:re,active:A.active}):ee=re;break;case"fast-forward":const q=this.fastForwardActive?n(Re,{clsPrefix:o},{default:()=>this.rtlEnabled?n(mt,null):n(wt,null)}):n(Re,{clsPrefix:o},{default:()=>n(xt,null)});X?ee=X({type:"fast-forward",node:q,active:this.fastForwardActive||this.showFastForwardMenu}):ee=q,be=this.handleFastForwardMouseenter,ge=this.handleFastForwardMouseleave;break;case"fast-backward":const j=this.fastBackwardActive?n(Re,{clsPrefix:o},{default:()=>this.rtlEnabled?n(wt,null):n(mt,null)}):n(Re,{clsPrefix:o},{default:()=>n(xt,null)});X?ee=X({type:"fast-backward",node:j,active:this.fastBackwardActive||this.showFastBackwardMenu}):ee=j,be=this.handleFastBackwardMouseenter,ge=this.handleFastBackwardMouseleave;break}const pe=n("div",{key:xe,class:[`${o}-pagination-item`,A.active&&`${o}-pagination-item--active`,Z!=="page"&&(Z==="fast-backward"&&this.showFastBackwardMenu||Z==="fast-forward"&&this.showFastForwardMenu)&&`${o}-pagination-item--hover`,s&&`${o}-pagination-item--disabled`,Z==="page"&&`${o}-pagination-item--clickable`],onClick:()=>ne(A),onMouseenter:be,onMouseleave:ge},ee);if(Z==="page"&&!A.mayBeFastBackward&&!A.mayBeFastForward)return pe;{const re=A.type==="page"?A.mayBeFastBackward?"fast-backward":"fast-forward":A.type;return n($n,{to:this.to,key:re,disabled:s,trigger:"hover",virtualScroll:!0,style:{width:"60px"},theme:h.peers.Popselect,themeOverrides:h.peerOverrides.Popselect,builtinThemeOverrides:{peers:{InternalSelectMenu:{height:"calc(var(--n-option-height) * 4.6)"}}},nodeProps:()=>({style:{justifyContent:"center"}}),show:Z==="page"?!1:Z==="fast-backward"?this.showFastBackwardMenu:this.showFastForwardMenu,onUpdateShow:q=>{Z!=="page"&&(q?Z==="fast-backward"?this.showFastBackwardMenu=q:this.showFastForwardMenu=q:(this.showFastBackwardMenu=!1,this.showFastForwardMenu=!1))},options:A.type!=="page"?A.options:[],onUpdateValue:this.handleMenuSelect,scrollable:!0,showCheckmark:!1},{default:()=>pe})}}),n("div",{class:[`${o}-pagination-item`,!K&&`${o}-pagination-item--button`,{[`${o}-pagination-item--disabled`]:f<1||f>=x||s}],onClick:se},K?K({page:f,pageSize:g,pageCount:x,itemCount:this.mergedItemCount,startIndex:this.startIndex,endIndex:this.endIndex}):n(Re,{clsPrefix:o},{default:()=>this.rtlEnabled?n(pt,null):n(Ct,null)})));case"size-picker":return!z&&i?n(En,Object.assign({consistentMenuWidth:!1,placeholder:"",showCheckmark:!1,to:this.to},this.selectProps,{size:v,options:c,value:g,disabled:s,theme:h.peers.Select,themeOverrides:h.peerOverrides.Select,onUpdateValue:N})):null;case"quick-jumper":return!z&&b?n("div",{class:`${o}-pagination-quick-jumper`},_?_():Sn(this.$slots.goto,()=>[F.goto]),n(ht,{value:p,onUpdateValue:L,size:y,placeholder:"",disabled:s,theme:h.peers.Input,themeOverrides:h.peerOverrides.Input,onChange:V})):null;default:return null}}),ve?n("div",{class:`${o}-pagination-suffix`},ve({page:f,pageSize:g,pageCount:x,startIndex:this.startIndex,endIndex:this.endIndex,itemCount:this.mergedItemCount})):null)}});export{Wn as _};
