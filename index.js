$(document).ready(function(){
    $("#relax").click(function(){
      $("#all").transition('slide left');
      setTimeout(function(){window.location.href="relax/relax.html";}, 500);
      
    });
});