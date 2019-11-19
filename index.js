$(document).ready(function(){
    starFild();
    terrain();
    tree();
    window.onresize = function(){
        starFild();
        terrain();
        tree();
     };
});