function terrain() {
	// Terrain stuff.
	var terrain = document.getElementById("terCanvas");
	var terCtx = terrain.getContext("2d");
	terCtx.clearRect(0, 0, terrain.width, terrain.height);
	var width = $("#canvas_space").parent().width();
	var height = $("#canvas_space").parent().height();
	terrain.width = width;
	terrain.height = height;

	// Some random points
	var points = [];
	var displacement = 50;
	var power = Math.pow(2, Math.ceil(Math.log(screen.width) / (Math.log(2))));

	//altura esquerda
	points[0] = height - height / 2;
	//direita
	points[power] = height;

	// create the rest of the points
	for (var i = 1; i < power; i *= 2) {
		for (var j = (power / i) / 2; j < power; j += power / i) {
			points[j] = ((points[j - (power / i) / 2] + points[j + (power / i) / 2]) / 2) + Math.floor(Math.random() * -displacement + displacement);
		}
		displacement *= 0.6;
	}
	// draw the terrain
	terCtx.beginPath();

	for (var i = 0; i <= screen.width; i++) {
		if (i === 0) {
			terCtx.moveTo(0, points[0]);
		} else if (points[i] !== undefined) {
			terCtx.lineTo(i, points[i]);
		}
	}

	terCtx.lineTo(width, terrain.height);
	terCtx.lineTo(0, terrain.height);
	terCtx.lineTo(0, points[0]);
    terCtx.fill();
    //draw my name
	/*var name = "Guilherme Solski Alves";
	terCtx.font = "50% Comic Sans MS";
	terCtx.fillStyle = "white";
	terCtx.fillText(name, 0, height - height / 4);*/
}

