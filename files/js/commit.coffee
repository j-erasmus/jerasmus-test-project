class Commit
  constructor: ->
    $('.files .diff-file').each ->
      new CommitFile(@)

@Commit = Commit