var c={exports:{}},y;function q(){return y||(y=1,function(m){var u={single_source_shortest_paths:function(o,t,e){var r={},i={};i[t]=0;var n=u.PriorityQueue.make();n.push(t,0);for(var f,p,s,h,_,v,a,d,l;!n.empty();){f=n.pop(),p=f.value,h=f.cost,_=o[p]||{};for(s in _)_.hasOwnProperty(s)&&(v=_[s],a=h+v,d=i[s],l=typeof i[s]>"u",(l||d>a)&&(i[s]=a,n.push(s,a),r[s]=p))}if(typeof e<"u"&&typeof i[e]>"u"){var k=["Could not find a path from ",t," to ",e,"."].join("");throw new Error(k)}return r},extract_shortest_path_from_predecessor_list:function(o,t){for(var e=[],r=t;r;)e.push(r),o[r],r=o[r];return e.reverse(),e},find_path:function(o,t,e){var r=u.single_source_shortest_paths(o,t,e);return u.extract_shortest_path_from_predecessor_list(r,e)},PriorityQueue:{make:function(o){var t=u.PriorityQueue,e={},r;o=o||{};for(r in t)t.hasOwnProperty(r)&&(e[r]=t[r]);return e.queue=[],e.sorter=o.sorter||t.default_sorter,e},default_sorter:function(o,t){return o.cost-t.cost},push:function(o,t){var e={value:o,cost:t};this.queue.push(e),this.queue.sort(this.sorter)},pop:function(){return this.queue.shift()},empty:function(){return this.queue.length===0}}};m.exports=u}(c)),c.exports}export{q as r};
