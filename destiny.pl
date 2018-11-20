#!/usr/bin/perl

use strict;
use warnings;

no warnings 'uninitialized';

use autodie;

use JSON::PP;
use HTTP::Tiny;

my $swagger_url = 'https://github.com/kubernetes/kubernetes/raw/release-1.10/api/openapi-spec/swagger.json';

my $doc = do {
  my $response = HTTP::Tiny->new->get($swagger_url);

  die "Failed to download swagger.json\n" unless $response->{success};

  decode_json($response->{content});
};

my ($out_file, @defs) = @ARGV;

my $str = "";

$str .= "package manifests\n\n";
$str .= "func init() {\n";
$str .= "\tpaths = map[string]map[string][][]string{\n";
for my $def (@defs) {
  $str .= qq(\t\t"$def": map[string][][]string{\n);
  my %merged;
  for my $found (find($doc, "#/definitions/$def")) {
    $merged{$found->{definition}} ||= [];

    push @{$merged{$found->{definition}}}, $found->{path};
  }
  for my $from (keys %merged) {
    $str .= qq(\t\t\t"$from": [][]string{\n);

    for my $path (@{$merged{$from}}) {
      $str .= qq(\t\t\t\t[]string{) . join(", ", map qq("$_"), @$path) . qq(},\n);
    }
    $str .= qq(\t\t\t},\n);
  }
  $str .= qq(\t\t},\n);
}
$str .= "\t}\n";
$str .= "}\n\n";

$str .= "func init() {\n";
$str .= "\ttypes = map[ktype]string{\n";
my %defs = %{$doc->{definitions}};
for my $def (sort keys %defs) {
  my $mappings = $defs{$def}{'x-kubernetes-group-version-kind'};
  next if !$mappings;

  for my $m (@$mappings) {
    $str .=  qq(\t\tktype{group: "$m->{group}", kind: "$m->{kind}", version: "$m->{version}"}: "$def",\n);
  }

  $str .= "\n";

}
$str .= "\t}\n";
$str .= "}\n";

open my $fh, '>', $out_file;
print $fh $str;
close $fh;

sub _path_for {
   my ($doc, $needle) = @_;

   return [] if $doc eq $needle;
   return () if !ref $doc;

   my @ret;

   my @possible_properties = keys %$doc;
   for my $pprop (sort @possible_properties) {
      if ($pprop eq '$ref' && $doc->{$pprop} eq $needle) {
         push @ret, [];
         next;
      }

      my @found = _path_for($doc->{$pprop}, $needle);

      $pprop = "" if $doc->{type} eq "array";
      for my $found (@found) {
         push @ret, [$pprop, @$found]
      }
   }

   return @ret;
}

sub _find {
   my ($doc, $needle) = @_;

   my @ret;
   my $def = $doc->{definitions};
   for my $pdef (sort keys %$def) {
      for my $path (_path_for($def->{$pdef}{properties}, $needle)) {
         push @ret, {
           definition => $pdef,
           path => $path,
         }
      }
   }

   return @ret
}

sub find {
   my ($doc, $needle) = @_;

   my @ret;
   my @outer = _find($doc, "$needle");

   push @ret, @outer;

   for my $outer (@outer) {
      my @related = find($doc, "#/definitions/$outer->{definition}");

      for my $related (@related) {
          push @ret, {
            definition => $related->{definition},
            path => [@{$related->{path}},@{$outer->{path}}],
          }
      }
   }


   return @ret;
}
