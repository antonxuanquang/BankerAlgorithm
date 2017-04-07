var gulp  = require('gulp')
var shell = require('gulp-shell')
var watch = require('gulp-watch')

gulp.task('compile', shell.task([
	'clear',
	'go run bankerAlgorithm.go test1.txt',
	'sleep 3',
	'clear',
	'go run bankerAlgorithm.go test2.txt',
	'sleep 3',
	'clear',
	'go run bankerAlgorithm.go test3.txt'
]))

gulp.task('watch', function() {
	gulp.watch('./*.go', ['compile']);
	gulp.watch('./*.h', ['compile']);
});

gulp.task('default' ,function() {
	gulp.watch('./*.go', ['compile']);
	gulp.watch('./*.h', ['compile']);
});
