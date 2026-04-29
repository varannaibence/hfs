class Node {
    has Str $.value is rw;
    has Node $.zero is rw;
    has Node $.one is rw;
}

class LZWBinTree {
    has Node $.root is rw;
    has Node $.current is rw;
    has Int $.depth is rw = 0;
    has Int $.max-depth is rw = 0;
    has Int $.leaf-sum is rw = 0;
    has Int $.leaf-count is rw = 0;
    has Num $.variance-sum is rw = 0e0;
    has Num $.average is rw = 0e0;
    has Num $.deviation is rw = 0e0;

    submethod BUILD() {
        $!root = Node.new(value => "/");
        $!current = $!root;
    }

    method insert-bit(Str $bit) {
        if $bit eq "0" {
            if !$!current.zero.defined {
                $!current.zero = Node.new(value => "0");
                $!current = $!root;
            } else {
                $!current = $!current.zero;
            }
            return;
        }

        if !$!current.one.defined {
            $!current.one = Node.new(value => "1");
            $!current = $!root;
        } else {
            $!current = $!current.one;
        }
    }

    method write-tree($out) {
        $!depth = 0;
        self!write-node($!root, $out);
    }

    method !write-node(Node $node, $out) {
        return unless $node.defined;

        $!depth++;
        self!write-node($node.one, $out) if $node.one.defined;
        $out.print("---" x $!depth);
        $out.say($node.value ~ "(" ~ ($!depth - 1) ~ ")");
        self!write-node($node.zero, $out) if $node.zero.defined;
        $!depth--;
    }

    method get-depth() {
        $!depth = 0;
        $!max-depth = 0;
        self!measure-depth($!root);
        return $!max-depth - 1;
    }

    method !measure-depth(Node $node) {
        return unless $node.defined;

        $!depth++;
        if $!depth > $!max-depth {
            $!max-depth = $!depth;
        }
        self!measure-depth($node.one) if $node.one.defined;
        self!measure-depth($node.zero) if $node.zero.defined;
        $!depth--;
    }

    method get-mean() {
        $!depth = 0;
        $!leaf-sum = 0;
        $!leaf-count = 0;
        self!measure-mean($!root);
        $!average = ($!leaf-sum / $!leaf-count).Num;
        return $!average;
    }

    method !measure-mean(Node $node) {
        return unless $node.defined;

        $!depth++;
        self!measure-mean($node.one) if $node.one.defined;
        self!measure-mean($node.zero) if $node.zero.defined;
        $!depth--;

        if !$node.one.defined && !$node.zero.defined {
            $!leaf-count++;
            $!leaf-sum += $!depth;
        }
    }

    method get-deviation() {
        $!average = self.get-mean();
        $!variance-sum = 0e0;
        $!depth = 0;
        $!leaf-count = 0;
        self!measure-deviation($!root);

        if $!leaf-count - 1 > 0 {
            $!deviation = sqrt($!variance-sum / ($!leaf-count - 1));
        } else {
            $!deviation = sqrt($!variance-sum);
        }

        return $!deviation;
    }

    method !measure-deviation(Node $node) {
        return unless $node.defined;

        $!depth++;
        self!measure-deviation($node.one) if $node.one.defined;
        self!measure-deviation($node.zero) if $node.zero.defined;
        $!depth--;

        if !$node.one.defined && !$node.zero.defined {
            $!leaf-count++;
            my $difference = $!depth - $!average;
            $!variance-sum += $difference * $difference;
        }
    }
}

sub usage() {
    say "Usage: lzwtree in_file -o out_file";
}

sub process-input(Blob $data, LZWBinTree $tree) {
    my $index = 0;
    while $index < $data.elems && $data[$index] != 0x0a {
        $index++;
    }
    $index++ if $index < $data.elems;

    my $comment-mode = False;
    while $index < $data.elems {
        my $value = $data[$index];

        if $value == 0x3e {
            $comment-mode = True;
            $index++;
            next;
        }
        if $value == 0x0a {
            $comment-mode = False;
            $index++;
            next;
        }
        if $comment-mode || $value == 0x4e {
            $index++;
            next;
        }

        my $byte = $value;
        for ^8 {
            if $byte +& 0x80 {
                $tree.insert-bit("1");
            } else {
                $tree.insert-bit("0");
            }
            $byte = (($byte +< 1) +& 0xff);
        }

        $index++;
    }
}

if @*ARGS.elems != 3 {
    usage();
    exit(-1);
}

my $input-file = @*ARGS[0];
if @*ARGS[1] ne "-o" {
    usage();
    exit(-2);
}

unless $input-file.IO.e {
    say $input-file ~ " nem letezik...";
    usage();
    exit(-3);
}

my $data = slurp $input-file, :bin;
my $tree = LZWBinTree.new();
process-input($data, $tree);

my $output-file = @*ARGS[2];
my $out = open $output-file, :w;
$tree.write-tree($out);
$out.say("depth = " ~ $tree.get-depth());
$out.say("mean = " ~ $tree.get-mean());
$out.say("var = " ~ $tree.get-deviation());
$out.close();
