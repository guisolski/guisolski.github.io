function readjustment(){
  
  var width = $(window).width();
  var height = $(window).height();

  if(width > 760 && height > 150){
    $("#addSeg").addClass("segment");
    $("#addSeg").addClass("scrolling");
    $("#addSeg").addClass("rem50");
  }else{
    $("#addSeg").removeClass("segment");
    $("#addSeg").removeClass("scrolling");
    $("#addSeg").addClass("rem50");
  }
}
readjustment(); 
$(document).ready(function () {
  
  $("#relax").click(function () {
    $("#all").transition('slide left');
    setTimeout(function () { window.location.href = "relax/relax.html"; }, 500);
  });  
  window.onresize = function () {
    readjustment(); 
  };
  $('.overlay')
  .visibility({
    type   : 'fixed',
    offset : 5 // give some space from top of screen
  })
;
});
