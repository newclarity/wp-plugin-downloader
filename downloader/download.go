package downloader

import (
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/newclarity/wp-plugin-downloader/only"
	"github.com/newclarity/wp-plugin-downloader/svn"
	"os"
)

type DownloadArgs struct {
	What      string
	Basedir   string
	SvnDomain string
}

func Download(args *DownloadArgs) (sts status.Status) {
	for range only.Once {
		sts = EnsureDir(args.Basedir)
		if is.Error(sts) {
			break
		}
		repo := svn.NewSvn(args.SvnDomain,args.Basedir)
		rev,sts := repo.GetLatestRevision()
		if is.Error(sts) {
			panic(sts.Message())
		}
		fmt.Printf("Most recent SVN revision: %s\n",rev)
		rev,sts := repo.GetLastSyncedRevision()
		if is.Error(sts) {
			panic(sts.Message())
		}
		fmt.Printf("Most recent SVN revision: %s\n",rev)


		/*
		if ( file_exists( "{$directory}.last-revision" ) ) {
			$last_revision = (int) file_get_contents( "{$directory}.last-revision" );
			echo "Last synced revision: {$last_revision}\r\n";
		} else {
			$last_revision = false;
			echo "You have not yet performed a successful sync. Settle in. This will take a while.\r\n";
		}

		*/
	}
	return sts
}

func EnsureDir(dir string) (sts status.Status){
	err := os.Mkdir(dir,0777)
	if err != nil {
		sts = status.Wrap(err,&status.Args{
			Message: fmt.Sprintf("unable to make directory '%s'",dir),
		})
	}
	return sts
}

/*
if ( file_exists( "{$directory}.last-revision" ) ) {
	$last_revision = (int) file_get_contents( "{$directory}.last-revision" );
	echo "Last synced revision: {$last_revision}\r\n";
} else {
	$last_revision = false;
	echo "You have not yet performed a successful sync. Settle in. This will take a while.\r\n";
}

$start_time = time();
$plugins = array();

if ( $last_revision != $svn_last_revision ) {
	if ( $last_revision ) {
		$changelog_url = sprintf( 'http://plugins.trac.wordpress.org/log/?verbose=on&mode=follow_copy&format=changelog&rev=%d&limit=%d', $svn_last_revision, $svn_last_revision - $last_revision );
		$changes = file_get_contents( $changelog_url );
		$now = date( 'Y-m-d-h-i-s-a' );
		file_put_contents( "{$download_dir}changelogs/{$now}-changes.log", $changes );
		preg_match_all( '#^' . "\t" . '*\* ([^/A-Z ]+)[ /].* \((added|modified|deleted|moved|copied)\)' . "\n" . '#m', $changes, $matches );
		$plugins = array_unique( $matches[1] );
	} else {
		$plugins = file_get_contents( 'http://svn.wp-plugins.org/' );
		preg_match_all( '#<li><a href="([^/]+)/">([^/]+)/</a></li>#', $plugins, $matches );
		$plugins = $matches[1];
	}

  $output_file = tempnam( "{$download_dir}/temp", 'wget-output-' );

	foreach ( $plugins as $plugin ) {

	  if ( is_file( "{$download_dir}missing/{$plugin}.missing" ) ) {
	    echo "Skipping {$plugin} - previously found missing.\r\n";
	    continue;
    }

		$plugin = urldecode( $plugin );
		echo "Updating {$plugin}";

		$output = null; $return = null;
		exec( $cmd = "{$wget_dir}/wget -o {$output_file} -np -O " . escapeshellarg( sprintf( $download, $plugin ) ) . ' ' . escapeshellarg( sprintf( $url, $plugin ) ) . ' > /dev/null', $output, $return );

		if ( $return === 0 && file_exists( sprintf( $download, $plugin ) ) ) {
			if ($type === 'all') {
				if ( file_exists( "{$download_dir}plugins/{$plugin}" ) )
					exec( 'rm -rf ' . escapeshellarg( "{$download_dir}plugins/{$plugin}" ) );

				exec( 'unzip -o -d ' . escapeshellarg( "{$download_dir}plugins" ) . ' ' . escapeshellarg( "{$download_dir}zips/{$plugin}.zip" ) );
				exec( 'rm -rf ' . escapeshellarg( "{$download_dir}zips.zip" ) );
			}
		} else {
      echo "... download failed, return code {$return}.";
      $output = file_get_contents( $output_file );
      preg_match( '#ERROR ([1-5][0-9]{2}): (.+)$#m', $output, $matches );
      if ( '404' == $matches[1] )
        file_put_contents( "{$download_dir}missing/{$plugin}.missing", true );
      echo " ERROR: {$matches[1]} {$matches[2]}";

		}
		echo "\r\n";
	}

  exec( 'rm ' . escapeshellarg( $output_file ) );

	if ( file_put_contents( "{$directory}.last-revision", $svn_last_revision ) )
		echo "[CLEANUP] Updated {$directory}.last-revision to {$svn_last_revision}\r\n";
	else
		echo "[ERROR] Could not update {$directory}.last-revision to {$svn_last_revision}\r\n";
}

$end_time = time();
$minutes = ( $end_time - $start_time ) / 60;
$seconds = ( $end_time - $start_time ) % 60;

echo "[SUCCESS] Done updating plugins!\r\n";
echo "It took " . number_format($minutes) . " minute" . ( $minutes == 1 ? '' : 's' ) . " and " . $seconds . " second" . ( $seconds == 1 ? '' : 's' ) . " to update ". count($plugins)  ." plugin" . ( count($plugins) == 1 ? '' : 's') . "\r\n";
echo "[DONE]\r\n";

*/
