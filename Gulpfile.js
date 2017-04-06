var gulp  = require('gulp')
var shell = require('gulp-shell')
var watch = require('gulp-watch')

gulp.task('compile', shell.task([
  'clear',
  'make clean',
  'make osp2a',
  'echo',
  'echo "Test case 1"',
  './osp2a "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet, possimus! 2421"',
  'echo',
  'echo "Test case 2"',
  './osp2a Lorem ipsum dolor "sit amet, consectetur" adipisicing elit. Amet, possimus! 2421',
  'echo',
  'echo "Test case 3"',
  './osp2a Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet, possimus! 2421',
  'echo',
  'echo "Test case 4"',
  './osp2a Lorem ipsum """""" 2421',
  'make osp2b',
  'echo',
  './osp2b '  
              + 'part2_test/*'
              // + 'part2_test/black_diagonal.txt '
              // + 'part2_test/black_horizontal.txt '
              // + 'part2_test/black_rdiagonal.txt '
              // + 'part2_test/black_vertical.txt '

              // + 'part2_test/red_diagonal.txt '
              // + 'part2_test/red_horizontal.txt '
              // + 'part2_test/red_rdiagonal.txt '
              // + 'part2_test/red_vertical.txt '

              // + 'part2_test/no_win1.txt '
              // + 'part2_test/no_win2.txt '
              // + 'part2_test/no_win3.txt '
              ,
]))

gulp.task('watch', function() {
    gulp.watch('./*.c', ['compile']);
    gulp.watch('./*.h', ['compile']);
});

gulp.task('default' ,function() {
    gulp.watch('./*.c', ['compile']);
    gulp.watch('./*.h', ['compile']);
});
