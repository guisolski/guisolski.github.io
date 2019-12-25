function readjustment(){
  
  var width = $("#all").width();
  var height = $("#all").height();
  console.log(width);
  console.log(height);
  if(width > 760 && height > 150){
    $("#addSeg").addClass("segment");
    $("#addSeg").addClass("relax");
  }else{
    $("#addSeg").removeClass("segment");
    $("#addSeg").removeClass("relax");
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
});