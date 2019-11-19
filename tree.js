
function tree(){
    var tree = document.getElementById("treeCanvas");
    var treeCtx = tree.getContext("2d");
    var width = $("#canvas_space").parent().width();
    var height = $("#canvas_space").parent().height();
    tree.width = width;
    tree.height = height;
    console.log( parseInt(height*0.01));
    branch(treeCtx, width-width/5, height, -90, parseInt(height*0.01));
}
function branch(context, x1, y1, angle, depth){
	var branchArmLength = random(0, 20);
	if (depth != 0){
		var x2 = x1 + (Math.cos(angle*(Math.PI/180.0)) * depth * branchArmLength);
		var y2 = y1 + (Math.sin(angle*(Math.PI/180.0)) * depth * branchArmLength);
		
		line(context, x1, y1, x2, y2, depth);
		branch(context, x2, y2, angle - random(15,20), depth - 1);
		branch(context, x2, y2, angle + random(15,20), depth - 1);
	}
	else{
	
	var x2 = x1 + (Math.cos(angle*(Math.PI/180.0)) * depth * branchArmLength);
	var y2 = y1 + (Math.sin(angle*(Math.PI/180.0)) * depth * branchArmLength);
	context.fillStyle = 'white';
	context.arc(x2, y2, random(0,10), 0, 2 * Math.PI, false);
	context.fill();
	}
}
function line(context, x1, y1, x2, y2, thickness){
	context.strokeStyle = "Brown"		
	context.lineWidth = thickness * 1.5;
	context.beginPath();
	context.moveTo(x1,y1);
	context.lineTo(x2, y2);
	context.closePath();
	context.stroke();
}
function random(min, max){
	return min + Math.floor(Math.random()*(max+1-min));
}
