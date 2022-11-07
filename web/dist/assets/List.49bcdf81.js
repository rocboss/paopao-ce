import{n as d,l as n,o as z,m as i,aA as y,aB as _,d as P,u as R,q as c,A as $,s as j,y as B,h as a,e as M,aC as I}from"./index.f480f018.js";var L=d([n("list",`
 --n-merged-border-color: var(--n-border-color);
 --n-merged-color: var(--n-color)
 font-size: var(--n-font-size);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 padding: 0;
 list-style-type: none;
 color: var(--n-text-color);
 background-color: var(--n-merged-color);
 `,[z("bordered",`
 border-radius: var(--n-border-radius);
 border: 1px solid var(--n-merged-border-color);
 `,[n("list-item",`
 padding: 12px 20px;
 `,[d("&:not(:last-child)",`
 border-bottom: 1px solid var(--n-merged-border-color);
 `)]),i("header, footer",`
 padding: 12px 20px;
 `,[d("&:not(:last-child)",`
 border-bottom: 1px solid var(--n-merged-border-color);
 `)])]),i("header, footer",`
 padding: 12px 0;
 box-sizing: border-box;
 transition: border-color .3s var(--n-bezier);
 `,[d("&:not(:last-child)",`
 border-bottom: 1px solid var(--n-merged-border-color);
 `)]),n("list-item",`
 padding: 12px 0; 
 box-sizing: border-box;
 display: flex;
 flex-wrap: nowrap;
 align-items: center;
 transition: border-color .3s var(--n-bezier);
 `,[i("prefix",`
 margin-right: 20px;
 flex: 0;
 `),i("suffix",`
 margin-left: 20px;
 flex: 0;
 `),i("main",`
 flex: 1;
 `),d("&:not(:last-child)",`
 border-bottom: 1px solid var(--n-merged-border-color);
 `)])]),y(n("list",`
 --merged-color: var(--n-color-modal);
 --merged-border-color: var(--n-border-color-modal);
 `)),_(n("list",`
 --merged-color: var(--n-color-popover);
 --merged-border-color: var(--n-border-color-popover);
 `))]);const O=Object.assign(Object.assign({},c.props),{size:{type:String,default:"medium"},bordered:{type:Boolean,default:!1}}),T=M("n-list");var k=P({name:"List",props:O,setup(r){const{mergedClsPrefixRef:o,inlineThemeDisabled:e}=R(r),s=c("List","-list",L,I,r,o);$(T,{mergedClsPrefixRef:o});const t=j(()=>{const{common:{cubicBezierEaseInOut:b},self:{fontSize:m,textColor:p,color:v,colorModal:f,colorPopover:u,borderColor:g,borderColorModal:x,borderColorPopover:h,borderRadius:C}}=s.value;return{"--n-font-size":m,"--n-bezier":b,"--n-text-color":p,"--n-color":v,"--n-border-radius":C,"--n-border-color":g,"--n-border-color-modal":x,"--n-border-color-popover":h,"--n-color-modal":f,"--n-color-popover":u}}),l=e?B("list",void 0,t,r):void 0;return{mergedClsPrefix:o,cssVars:e?void 0:t,themeClass:l==null?void 0:l.themeClass,onRender:l==null?void 0:l.onRender}},render(){var r;const{$slots:o,mergedClsPrefix:e,onRender:s}=this;return s==null||s(),a("ul",{class:[`${e}-list`,this.bordered&&`${e}-list--bordered`,this.themeClass],style:this.cssVars},o.header?a("div",{class:`${e}-list__header`},o.header()):null,(r=o.default)===null||r===void 0?void 0:r.call(o),o.footer?a("div",{class:`${e}-list__footer`},o.footer()):null)}});export{k as _,T as l};
