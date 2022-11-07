import{d as O,h as a,aD as Oe,aE as Be,a1 as ce,b1 as Fe,s as w,r as F,j as q,av as Y,au as D,b2 as ue,b3 as te,aH as J,b4 as Le,l as M,m as L,n as U,u as fe,q as K,b5 as $e,I as Ee,i as Z,b6 as He,v as H,y as ve,N as he,b7 as Q,aX as Ve,aJ as me,o as A,p as je,aO as Ae,t as B,b8 as De,b9 as Ke,S as ne,aF as Ge,ba as oe,A as ie,bb as We,z as qe,ay as Ue,bc as Ze,K as Xe,bd as Je,a$ as le,be as Ye,c as Qe,aa as et}from"./index.f480f018.js";import{l as tt}from"./List.49bcdf81.js";let re=!1;function nt(){if(typeof window!="undefined"&&!!window.CSS&&!re&&(re=!0,"registerProperty"in(window==null?void 0:window.CSS)))try{CSS.registerProperty({name:"--n-color-start",syntax:"<color>",inherits:!1,initialValue:"#0000"}),CSS.registerProperty({name:"--n-color-end",syntax:"<color>",inherits:!1,initialValue:"#0000"})}catch{}}var ot=O({name:"Checkmark",render(){return a("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 16 16"},a("g",{fill:"none"},a("path",{d:"M14.046 3.486a.75.75 0 0 1-.032 1.06l-7.93 7.474a.85.85 0 0 1-1.188-.022l-2.68-2.72a.75.75 0 1 1 1.068-1.053l2.234 2.267l7.468-7.038a.75.75 0 0 1 1.06.032z",fill:"currentColor"})))}}),it=O({name:"Empty",render(){return a("svg",{viewBox:"0 0 28 28",fill:"none",xmlns:"http://www.w3.org/2000/svg"},a("path",{d:"M26 7.5C26 11.0899 23.0899 14 19.5 14C15.9101 14 13 11.0899 13 7.5C13 3.91015 15.9101 1 19.5 1C23.0899 1 26 3.91015 26 7.5ZM16.8536 4.14645C16.6583 3.95118 16.3417 3.95118 16.1464 4.14645C15.9512 4.34171 15.9512 4.65829 16.1464 4.85355L18.7929 7.5L16.1464 10.1464C15.9512 10.3417 15.9512 10.6583 16.1464 10.8536C16.3417 11.0488 16.6583 11.0488 16.8536 10.8536L19.5 8.20711L22.1464 10.8536C22.3417 11.0488 22.6583 11.0488 22.8536 10.8536C23.0488 10.6583 23.0488 10.3417 22.8536 10.1464L20.2071 7.5L22.8536 4.85355C23.0488 4.65829 23.0488 4.34171 22.8536 4.14645C22.6583 3.95118 22.3417 3.95118 22.1464 4.14645L19.5 6.79289L16.8536 4.14645Z",fill:"currentColor"}),a("path",{d:"M25 22.75V12.5991C24.5572 13.0765 24.053 13.4961 23.5 13.8454V16H17.5L17.3982 16.0068C17.0322 16.0565 16.75 16.3703 16.75 16.75C16.75 18.2688 15.5188 19.5 14 19.5C12.4812 19.5 11.25 18.2688 11.25 16.75L11.2432 16.6482C11.1935 16.2822 10.8797 16 10.5 16H4.5V7.25C4.5 6.2835 5.2835 5.5 6.25 5.5H12.2696C12.4146 4.97463 12.6153 4.47237 12.865 4H6.25C4.45507 4 3 5.45507 3 7.25V22.75C3 24.5449 4.45507 26 6.25 26H21.75C23.5449 26 25 24.5449 25 22.75ZM4.5 22.75V17.5H9.81597L9.85751 17.7041C10.2905 19.5919 11.9808 21 14 21L14.215 20.9947C16.2095 20.8953 17.842 19.4209 18.184 17.5H23.5V22.75C23.5 23.7165 22.7165 24.5 21.75 24.5H6.25C5.2835 24.5 4.5 23.7165 4.5 22.75Z",fill:"currentColor"}))}}),lt=O({props:{onFocus:Function,onBlur:Function},setup(e){return()=>a("div",{style:"width: 0; height: 0",tabindex:0,onFocus:e.onFocus,onBlur:e.onBlur})}});function se(e){return e&-e}class rt{constructor(t,o){this.l=t,this.min=o;const i=new Array(t+1);for(let r=0;r<t+1;++r)i[r]=0;this.ft=i}add(t,o){if(o===0)return;const{l:i,ft:r}=this;for(t+=1;t<=i;)r[t]+=o,t+=se(t)}get(t){return this.sum(t+1)-this.sum(t)}sum(t){if(t===0)return 0;const{ft:o,min:i,l:r}=this;if(t===void 0&&(t=r),t>r)throw new Error("[FinweckTree.sum]: `i` is larger than length.");let d=t*i;for(;t>0;)d+=o[t],t-=se(t);return d}getBound(t){let o=0,i=this.l;for(;i>o;){const r=Math.floor((o+i)/2),d=this.sum(r);if(d>t){i=r;continue}else if(d<t){if(o===r)return this.sum(o+1)<=t?o+1:r;o=r}else return r}return o}}const st=J(".v-vl",{maxHeight:"inherit",height:"100%",overflow:"auto",minWidth:"1px"},[J("&:not(.v-vl--show-scrollbar)",{scrollbarWidth:"none"},[J("&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb",{width:0,height:0,display:"none"})])]);var at=O({name:"VirtualList",inheritAttrs:!1,props:{showScrollbar:{type:Boolean,default:!0},items:{type:Array,default:()=>[]},itemSize:{type:Number,required:!0},itemResizable:Boolean,itemsStyle:[String,Object],visibleItemsTag:{type:[String,Object],default:"div"},visibleItemsProps:Object,ignoreItemResize:Boolean,onScroll:Function,onWheel:Function,onResize:Function,defaultScrollKey:[Number,String],defaultScrollIndex:Number,keyField:{type:String,default:"key"},paddingTop:{type:[Number,String],default:0},paddingBottom:{type:[Number,String],default:0}},setup(e){const t=Oe();st.mount({id:"vueuc/virtual-list",head:!0,anchorMetaName:Be,ssr:t}),ce(()=>{const{defaultScrollIndex:s,defaultScrollKey:c}=e;s!=null?m({index:s}):c!=null&&m({key:c})}),Fe(()=>{m({top:k.value})});const o=w(()=>{const s=new Map,{keyField:c}=e;return e.items.forEach((R,S)=>{s.set(R[c],S)}),s}),i=F(null),r=F(void 0),d=new Map,u=w(()=>{const{items:s,itemSize:c,keyField:R}=e,S=new rt(s.length,c);return s.forEach((C,z)=>{const b=C[R],_=d.get(b);_!==void 0&&S.add(z,_)}),S}),h=F(0),k=F(0),f=q(()=>Math.max(u.value.getBound(k.value-Y(e.paddingTop))-1,0)),v=w(()=>{const{value:s}=r;if(s===void 0)return[];const{items:c,itemSize:R}=e,S=f.value,C=Math.min(S+Math.ceil(s/R+1),c.length-1),z=[];for(let b=S;b<=C;++b)z.push(c[b]);return z}),m=s=>{const{left:c,top:R,index:S,key:C,position:z,behavior:b,debounce:_=!0}=s;if(c!==void 0||R!==void 0)g(c,R,b);else if(S!==void 0)y(S,b,_);else if(C!==void 0){const V=o.value.get(C);V!==void 0&&y(V,b,_)}else z==="bottom"?g(0,Number.MAX_SAFE_INTEGER,b):z==="top"&&g(0,0,b)};function y(s,c,R){const{value:S}=u,C=S.sum(s)+Y(e.paddingTop);if(!R)i.value.scrollTo({left:0,top:C,behavior:c});else{const{scrollTop:z,offsetHeight:b}=i.value;if(C>z){const _=S.get(s);C+_<=z+b||i.value.scrollTo({left:0,top:C+_-b,behavior:c})}else i.value.scrollTo({left:0,top:C,behavior:c})}I=s}function g(s,c,R){i.value.scrollTo({left:s,top:c,behavior:R})}function p(s,c){var R,S,C,z;if(e.ignoreItemResize||E(c.target))return;const{value:b}=u,_=o.value.get(s),V=b.get(_),j=(C=(S=(R=c.borderBoxSize)===null||R===void 0?void 0:R[0])===null||S===void 0?void 0:S.blockSize)!==null&&C!==void 0?C:c.contentRect.height;if(j===V)return;j-e.itemSize===0?d.delete(s):d.set(s,j-e.itemSize);const P=j-V;P!==0&&($!==void 0&&_<=$&&((z=i.value)===null||z===void 0||z.scrollBy(0,P)),b.add(_,P),h.value++)}function x(s){Le(G);const{onScroll:c}=e;c!==void 0&&c(s)}function N(s){if(E(s.target)||s.contentRect.height===r.value)return;r.value=s.contentRect.height;const{onResize:c}=e;c!==void 0&&c(s)}let I,$;function G(){const{value:s}=i;s!=null&&($=I!=null?I:f.value,I=void 0,k.value=i.value.scrollTop)}function E(s){let c=s;for(;c!==null;){if(c.style.display==="none")return!0;c=c.parentElement}return!1}return{listHeight:r,listStyle:{overflow:"auto"},keyToIndex:o,itemsStyle:w(()=>{const{itemResizable:s}=e,c=D(u.value.sum());return h.value,[e.itemsStyle,{boxSizing:"content-box",height:s?"":c,minHeight:s?c:"",paddingTop:D(e.paddingTop),paddingBottom:D(e.paddingBottom)}]}),visibleItemsStyle:w(()=>(h.value,{transform:`translateY(${D(u.value.sum(f.value))})`})),viewportItems:v,listElRef:i,itemsElRef:F(null),scrollTo:m,handleListResize:N,handleListScroll:x,handleItemResize:p}},render(){const{itemResizable:e,keyField:t,keyToIndex:o,visibleItemsTag:i}=this;return a(te,{onResize:this.handleListResize},{default:()=>{var r,d;return a("div",ue(this.$attrs,{class:["v-vl",this.showScrollbar&&"v-vl--show-scrollbar"],onScroll:this.handleListScroll,onWheel:this.onWheel,ref:"listElRef"}),[this.items.length!==0?a("div",{ref:"itemsElRef",class:"v-vl-items",style:this.itemsStyle},[a(i,Object.assign({class:"v-vl-visible-items",style:this.visibleItemsStyle},this.visibleItemsProps),{default:()=>this.viewportItems.map(u=>{const h=u[t],k=o.get(h),f=this.$slots.default({item:u,index:k})[0];return e?a(te,{key:h,onResize:v=>this.handleItemResize(h,v)},{default:()=>f}):(f.key=h,f)})})]):(d=(r=this.$slots).empty)===null||d===void 0?void 0:d.call(r)])}})}}),dt=M("empty",`
 display: flex;
 flex-direction: column;
 align-items: center;
 font-size: var(--n-font-size);
`,[L("icon",`
 width: var(--n-icon-size);
 height: var(--n-icon-size);
 font-size: var(--n-icon-size);
 line-height: var(--n-icon-size);
 color: var(--n-icon-color);
 transition:
 color .3s var(--n-bezier);
 `,[U("+",[L("description",`
 margin-top: 8px;
 `)])]),L("description",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 `),L("extra",`
 text-align: center;
 transition: color .3s var(--n-bezier);
 margin-top: 12px;
 color: var(--n-extra-text-color);
 `)]);const ct=Object.assign(Object.assign({},K.props),{description:String,showDescription:{type:Boolean,default:!0},showIcon:{type:Boolean,default:!0},size:{type:String,default:"medium"},renderIcon:Function});var ut=O({name:"Empty",props:ct,setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:o}=fe(e),i=K("Empty","-empty",dt,$e,e,t),{localeRef:r}=Ee("Empty"),d=Z(He,null),u=w(()=>{var v,m,y;return(v=e.description)!==null&&v!==void 0?v:(y=(m=d==null?void 0:d.mergedComponentPropsRef.value)===null||m===void 0?void 0:m.Empty)===null||y===void 0?void 0:y.description}),h=w(()=>{var v,m;return((m=(v=d==null?void 0:d.mergedComponentPropsRef.value)===null||v===void 0?void 0:v.Empty)===null||m===void 0?void 0:m.renderIcon)||(()=>a(it,null))}),k=w(()=>{const{size:v}=e,{common:{cubicBezierEaseInOut:m},self:{[H("iconSize",v)]:y,[H("fontSize",v)]:g,textColor:p,iconColor:x,extraTextColor:N}}=i.value;return{"--n-icon-size":y,"--n-font-size":g,"--n-bezier":m,"--n-text-color":p,"--n-icon-color":x,"--n-extra-text-color":N}}),f=o?ve("empty",w(()=>{let v="";const{size:m}=e;return v+=m[0],v}),k,e):void 0;return{mergedClsPrefix:t,mergedRenderIcon:h,localizedDescription:w(()=>u.value||r.value.description),cssVars:o?void 0:k,themeClass:f==null?void 0:f.themeClass,onRender:f==null?void 0:f.onRender}},render(){const{$slots:e,mergedClsPrefix:t,onRender:o}=this;return o==null||o(),a("div",{class:[`${t}-empty`,this.themeClass],style:this.cssVars},this.showIcon?a("div",{class:`${t}-empty__icon`},e.icon?e.icon():a(he,{clsPrefix:t},{default:this.mergedRenderIcon})):null,this.showDescription?a("div",{class:`${t}-empty__description`},e.default?e.default():this.localizedDescription):null,e.extra?a("div",{class:`${t}-empty__extra`},e.extra()):null)}});const ft=a(ot);function vt(e,t){return a(Ve,{name:"fade-in-scale-up-transition"},{default:()=>e?a(he,{clsPrefix:t,class:`${t}-base-select-option__check`},{default:()=>ft}):null})}var ae=O({name:"NBaseSelectOption",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(e){const{valueRef:t,pendingTmNodeRef:o,multipleRef:i,valueSetRef:r,renderLabelRef:d,renderOptionRef:u,handleOptionClick:h,handleOptionMouseEnter:k}=Z(Q),f=q(()=>{const{value:g}=o;return g?e.tmNode.key===g.key:!1});function v(g){const{tmNode:p}=e;p.disabled||h(g,p)}function m(g){const{tmNode:p}=e;p.disabled||k(g,p)}function y(g){const{tmNode:p}=e,{value:x}=f;p.disabled||x||k(g,p)}return{multiple:i,isGrouped:q(()=>{const{tmNode:g}=e,{parent:p}=g;return p&&p.rawNode.type==="group"}),isPending:f,isSelected:q(()=>{const{value:g}=t,{value:p}=i;if(g===null)return!1;const x=e.tmNode.rawNode.value;if(p){const{value:N}=r;return N.has(x)}else return g===x}),renderLabel:d,renderOption:u,handleMouseMove:y,handleMouseEnter:m,handleClick:v}},render(){const{clsPrefix:e,tmNode:{rawNode:t},isSelected:o,isPending:i,isGrouped:r,multiple:d,renderOption:u,renderLabel:h,handleClick:k,handleMouseEnter:f,handleMouseMove:v}=this,y=vt(d&&o,e),g=h?[h(t,o),y]:[me(t.label,t,o),y],p=a("div",{class:[`${e}-base-select-option`,t.class,{[`${e}-base-select-option--disabled`]:t.disabled,[`${e}-base-select-option--selected`]:o,[`${e}-base-select-option--grouped`]:r,[`${e}-base-select-option--pending`]:i}],style:t.style,onClick:k,onMouseenter:f,onMousemove:v},a("div",{class:`${e}-base-select-option__content`},g));return t.render?t.render({node:p,option:t,selected:o}):u?u({node:p,option:t,selected:o}):p}}),de=O({name:"NBaseSelectGroupHeader",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(){const{renderLabelRef:e,renderOptionRef:t}=Z(Q);return{renderLabel:e,renderOption:t}},render(){const{clsPrefix:e,renderLabel:t,renderOption:o,tmNode:{rawNode:i}}=this,r=t?t(i,!1):me(i.label,i,!1),d=a("div",{class:`${e}-base-select-group-header`},r);return i.render?i.render({node:d,option:i}):o?o({node:d,option:i,selected:!1}):d}}),ht=M("base-select-menu",`
 line-height: 1.5;
 outline: none;
 z-index: 0;
 position: relative;
 border-radius: var(--n-border-radius);
 transition:
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-color);
`,[A("multiple",[M("base-select-option",`
 padding-right: 28px;
 `)]),M("scrollbar",`
 max-height: var(--n-height);
 `),M("virtual-list",`
 max-height: var(--n-height);
 `),M("base-select-option",`
 min-height: var(--n-option-height);
 font-size: var(--n-option-font-size);
 display: flex;
 align-items: center;
 `,[L("content",`
 white-space: nowrap;
 text-overflow: ellipsis;
 overflow: hidden;
 `)]),M("base-select-group-header",`
 min-height: var(--n-option-height);
 font-size: .93em;
 display: flex;
 align-items: center;
 `),M("base-select-menu-option-wrapper",`
 position: relative;
 width: 100%;
 `),L("loading, empty",`
 display: flex;
 padding: 12px 32px;
 flex: 1;
 justify-content: center;
 `),L("loading",`
 color: var(--n-loading-color);
 font-size: var(--n-loading-size);
 `),L("action",`
 padding: 8px var(--n-option-padding-left);
 font-size: var(--n-option-font-size);
 transition: 
 color .3s var(--n-bezier);
 border-color .3s var(--n-bezier);
 border-top: 1px solid var(--n-action-divider-color);
 color: var(--n-action-text-color);
 `),M("base-select-group-header",`
 position: relative;
 cursor: default;
 padding: var(--n-option-padding);
 color: var(--n-group-header-text-color);
 `),M("base-select-option",`
 cursor: pointer;
 position: relative;
 padding: var(--n-option-padding);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 box-sizing: border-box;
 color: var(--n-option-text-color);
 opacity: 1;
 `,[U("&:active",`
 color: var(--n-option-text-color-pressed);
 `),A("grouped",`
 padding-left: calc(var(--n-option-padding-left) * 1.5);
 `),A("pending",`
 background-color: var(--n-option-color-pending);
 `),A("selected",`
 color: var(--n-option-text-color-active);
 background-color: var(--n-option-color-active);
 `),A("disabled",`
 cursor: not-allowed;
 `,[je("selected",`
 color: var(--n-option-text-color-disabled);
 `),A("selected",`
 opacity: var(--n-option-opacity-disabled);
 `)]),L("check",`
 font-size: 16px;
 position: absolute;
 right: 8px;
 top: calc(50% - 7px);
 color: var(--n-option-check-color);
 transition: color .3s var(--n-bezier);
 `,[Ae({enterScale:"0.5"})])])]),St=O({name:"InternalSelectMenu",props:Object.assign(Object.assign({},K.props),{clsPrefix:{type:String,required:!0},scrollable:{type:Boolean,default:!0},treeMate:{type:Object,required:!0},multiple:Boolean,size:{type:String,default:"medium"},value:{type:[String,Number,Array],default:null},width:[Number,String],autoPending:Boolean,virtualScroll:{type:Boolean,default:!0},show:{type:Boolean,default:!0},loading:Boolean,focusable:Boolean,renderLabel:Function,renderOption:Function,onMousedown:Function,onScroll:Function,onFocus:Function,onBlur:Function,onKeyup:Function,onKeydown:Function,onTabOut:Function,onMouseenter:Function,onMouseleave:Function,resetMenuOnOptionsChange:{type:Boolean,default:!0},inlineThemeDisabled:Boolean,onToggle:Function}),setup(e){const t=K("InternalSelectMenu","-internal-select-menu",ht,De,e,B(e,"clsPrefix")),o=F(null),i=F(null),r=F(null),d=w(()=>e.treeMate.getFlattenedNodes()),u=w(()=>Ke(d.value)),h=F(null);function k(){const{treeMate:n}=e;let l=null;if(e.autoPending){const{value:T}=e;T===null?l=n.getFirstAvailableNode():(e.multiple?l=n.getNode((T||[])[(T||[]).length-1]):l=n.getNode(T),(!l||l.disabled)&&(l=n.getFirstAvailableNode())),l&&z(l)}}let f;ne(B(e,"show"),n=>{n?f=ne(e.resetMenuOnOptionsChange?[B(e,"treeMate"),B(e,"multiple")]:[B(e,"multiple")],()=>{k(),Ge(b)},{immediate:!0}):f==null||f()},{immediate:!0});const v=w(()=>Y(t.value.self[H("optionHeight",e.size)])),m=w(()=>oe(t.value.self[H("padding",e.size)])),y=w(()=>e.multiple&&Array.isArray(e.value)?new Set(e.value):new Set),g=w(()=>{const n=d.value;return n&&n.length===0});function p(n){const{onToggle:l}=e;l&&l(n)}function x(n){const{onScroll:l}=e;l&&l(n)}function N(n){var l;(l=r.value)===null||l===void 0||l.sync(),x(n)}function I(){var n;(n=r.value)===null||n===void 0||n.sync()}function $(){const{value:n}=h;return n||null}function G(n,l){l.disabled||z(l,!1)}function E(n,l){l.disabled||p(l)}function s(n){var l;le(n,"action")||(l=e.onKeyup)===null||l===void 0||l.call(e,n)}function c(n){var l;le(n,"action")||(l=e.onKeydown)===null||l===void 0||l.call(e,n)}function R(n){var l;(l=e.onMousedown)===null||l===void 0||l.call(e,n),!e.focusable&&n.preventDefault()}function S(){const{value:n}=h;n&&z(n.getNext({loop:!0}),!0)}function C(){const{value:n}=h;n&&z(n.getPrev({loop:!0}),!0)}function z(n,l=!1){h.value=n,l&&b()}function b(){var n,l;const T=h.value;if(!T)return;const W=u.value(T.key);W!==null&&(e.virtualScroll?(n=i.value)===null||n===void 0||n.scrollTo({index:W}):(l=r.value)===null||l===void 0||l.scrollTo({index:W,elSize:v.value}))}function _(n){var l,T;!((l=o.value)===null||l===void 0)&&l.contains(n.target)&&((T=e.onFocus)===null||T===void 0||T.call(e,n))}function V(n){var l,T;!((l=o.value)===null||l===void 0)&&l.contains(n.relatedTarget)||(T=e.onBlur)===null||T===void 0||T.call(e,n)}ie(Q,{handleOptionMouseEnter:G,handleOptionClick:E,valueSetRef:y,multipleRef:B(e,"multiple"),valueRef:B(e,"value"),renderLabelRef:B(e,"renderLabel"),renderOptionRef:B(e,"renderOption"),pendingTmNodeRef:h}),ie(We,o),ce(()=>{const{value:n}=r;n&&n.sync()});const j=w(()=>{const{size:n}=e,{common:{cubicBezierEaseInOut:l},self:{height:T,borderRadius:W,color:pe,groupHeaderTextColor:be,actionDividerColor:ye,optionTextColorPressed:xe,optionTextColor:Se,optionTextColorDisabled:ze,optionTextColorActive:we,optionOpacityDisabled:Re,optionCheckColor:Ce,actionTextColor:ke,optionColorPending:Te,optionColorActive:_e,loadingColor:Ne,loadingSize:Me,[H("optionFontSize",n)]:Ie,[H("optionHeight",n)]:Pe,[H("optionPadding",n)]:ee}}=t.value;return{"--n-height":T,"--n-action-divider-color":ye,"--n-action-text-color":ke,"--n-bezier":l,"--n-border-radius":W,"--n-color":pe,"--n-option-font-size":Ie,"--n-group-header-text-color":be,"--n-option-check-color":Ce,"--n-option-color-pending":Te,"--n-option-color-active":_e,"--n-option-height":Pe,"--n-option-opacity-disabled":Re,"--n-option-text-color":Se,"--n-option-text-color-active":we,"--n-option-text-color-disabled":ze,"--n-option-text-color-pressed":xe,"--n-option-padding":ee,"--n-option-padding-left":oe(ee,"left"),"--n-loading-color":Ne,"--n-loading-size":Me}}),{inlineThemeDisabled:X}=e,P=X?ve("internal-select-menu",w(()=>e.size[0]),j,e):void 0,ge={selfRef:o,next:S,prev:C,getPendingTmNode:$};return Object.assign({mergedTheme:t,virtualListRef:i,scrollbarRef:r,itemSize:v,padding:m,flattenedNodes:d,empty:g,virtualListContainer(){const{value:n}=i;return n==null?void 0:n.listElRef},virtualListContent(){const{value:n}=i;return n==null?void 0:n.itemsElRef},doScroll:x,handleFocusin:_,handleFocusout:V,handleKeyUp:s,handleKeyDown:c,handleMouseDown:R,handleVirtualListResize:I,handleVirtualListScroll:N,cssVars:X?void 0:j,themeClass:P==null?void 0:P.themeClass,onRender:P==null?void 0:P.onRender},ge)},render(){const{$slots:e,virtualScroll:t,clsPrefix:o,mergedTheme:i,themeClass:r,onRender:d}=this;return d==null||d(),a("div",{ref:"selfRef",tabindex:this.focusable?0:-1,class:[`${o}-base-select-menu`,r,this.multiple&&`${o}-base-select-menu--multiple`],style:[{width:Je(this.width)},this.cssVars],onFocusin:this.handleFocusin,onFocusout:this.handleFocusout,onKeyup:this.handleKeyUp,onKeydown:this.handleKeyDown,onMousedown:this.handleMouseDown,onMouseenter:this.onMouseenter,onMouseleave:this.onMouseleave},this.loading?a("div",{class:`${o}-base-select-menu__loading`},a(Ue,{clsPrefix:o,strokeWidth:20})):this.empty?a("div",{class:`${o}-base-select-menu__empty`},Xe(e.empty,()=>[a(ut,{theme:i.peers.Empty,themeOverrides:i.peerOverrides.Empty})])):a(Ze,{ref:"scrollbarRef",theme:i.peers.Scrollbar,themeOverrides:i.peerOverrides.Scrollbar,scrollable:this.scrollable,container:t?this.virtualListContainer:void 0,content:t?this.virtualListContent:void 0,onScroll:t?void 0:this.doScroll},{default:()=>t?a(at,{ref:"virtualListRef",class:`${o}-virtual-list`,items:this.flattenedNodes,itemSize:this.itemSize,showScrollbar:!1,paddingTop:this.padding.top,paddingBottom:this.padding.bottom,onResize:this.handleVirtualListResize,onScroll:this.handleVirtualListScroll,itemResizable:!0},{default:({item:u})=>u.isGroup?a(de,{key:u.key,clsPrefix:o,tmNode:u}):u.ignored?null:a(ae,{clsPrefix:o,key:u.key,tmNode:u})}):a("div",{class:`${o}-base-select-menu-option-wrapper`,style:{paddingTop:this.padding.top,paddingBottom:this.padding.bottom}},this.flattenedNodes.map(u=>u.isGroup?a(de,{key:u.key,clsPrefix:o,tmNode:u}):a(ae,{clsPrefix:o,key:u.key,tmNode:u})))}),qe(e.action,u=>u&&[a("div",{class:`${o}-base-select-menu__action`,"data-action":!0,key:"action"},u),a(lt,{onFocus:this.onTabOut,key:"focus-detector"})]))}}),zt=O({name:"ListItem",setup(){const e=Z(tt,null);return e||Ye("list-item","`n-list-item` must be placed in `n-list`."),{mergedClsPrefix:e.mergedClsPrefixRef}},render(){const{$slots:e,mergedClsPrefix:t}=this;return a("li",{class:`${t}-list-item`},e.prefix?a("div",{class:`${t}-list-item__prefix`},e.prefix()):null,e.default?a("div",{class:`${t}-list-item__main`},e):null,e.suffix?a("div",{class:`${t}-list-item__suffix`},e.suffix()):null)}});const mt=e=>{const{heightSmall:t,heightMedium:o,heightLarge:i,borderRadius:r}=e;return{color:"#eee",colorEnd:"#ddd",borderRadius:r,heightSmall:t,heightMedium:o,heightLarge:i}},gt={name:"Skeleton",common:Qe,self:mt};var pt=U([M("skeleton",`
 height: 1em;
 width: 100%;
 transition: background-color .3s var(--n-bezier);
 transition:
 --n-color-start .3s var(--n-bezier),
 --n-color-end .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 animation: 2s skeleton-loading infinite cubic-bezier(0.36, 0, 0.64, 1);
 background-color: var(--n-color-start);
 `),U("@keyframes skeleton-loading",`
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
 `)]);const bt=Object.assign(Object.assign({},K.props),{text:Boolean,round:Boolean,circle:Boolean,height:[String,Number],width:[String,Number],size:String,repeat:{type:Number,default:1},animated:{type:Boolean,default:!0},sharp:{type:Boolean,default:!0}});var wt=O({name:"Skeleton",inheritAttrs:!1,props:bt,setup(e){nt();const{mergedClsPrefixRef:t}=fe(e),o=K("Skeleton","-skeleton",pt,gt,e,t);return{mergedClsPrefix:t,style:w(()=>{var i,r;const d=o.value,{common:{cubicBezierEaseInOut:u}}=d,h=d.self,{color:k,colorEnd:f,borderRadius:v}=h;let m;const{circle:y,sharp:g,round:p,width:x,height:N,size:I,text:$,animated:G}=e;I!==void 0&&(m=h[H("height",I)]);const E=y?(i=x!=null?x:N)!==null&&i!==void 0?i:m:x,s=(r=y&&x!=null?x:N)!==null&&r!==void 0?r:m;return{display:$?"inline-block":"",verticalAlign:$?"-0.125em":"",borderRadius:y?"50%":p?"4096px":g?"":v,width:typeof E=="number"?D(E):E,height:typeof s=="number"?D(s):s,animation:G?"":"none","--n-bezier":u,"--n-color-start":k,"--n-color-end":f}})}},render(){const{repeat:e,style:t,mergedClsPrefix:o,$attrs:i}=this,r=a("div",ue({class:`${o}-skeleton`,style:t},i));return e>1?a(et,null,Array.apply(null,{length:e}).map(d=>[r,`
`])):r}});export{St as N,zt as _,ut as a,wt as b,nt as u};
