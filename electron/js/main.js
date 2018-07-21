$(function() {
    $('.js-change').change(function() {
	var config = getValues();
	$.ajax("http://jfh.wtf:3720/settings?id=3",{
	    method: "POST",
	    data: JSON.stringify(config),
	    contentType: "application/json"
	});
    });
    
});

function getValues() {
    return {
	Orbit: parseInt($('#orbit').val()),
	Zoom: parseInt($('#zoom').val()),
	PresetValue: parseInt($("input[name='simpleSelection']:checked").val()),
	DisplayMode: parseInt($("input[name='displayMode']:checked").val()),
	View: parseInt($("input[name='view']:checked").val()),
	LayerTree: parseInt($("#tree-layer").val()),
	LayerCirculation: parseInt($("#circulation-layer").val()),
	LayerSiteBuilding: parseInt($("#sitebuilding-layer").val()),
	LayerProjectBuilding: parseInt($("#projectbuilding-layer").val()),
	LayerSunShadow: parseInt($("#sunshadow-layer").val()),
	LayerRoad: parseInt($("#road-layer").val())
    }
}
