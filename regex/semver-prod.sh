(?nx)^
(?<Major>0|[1-9]\d*)\.
(?<Minor>0|[1-9]\d*)\.
(?<Patch>0|[1-9]\d*)
(?<PreReleaseTagWithSeparator>
  -(?<PreReleaseTag>
    ((0|[1-9]\d*|\d*[A-Z-a-z-][\dA-Za-z-]*))(\.(0|[1-9]\d*|\d*[A-Za-z-][\dA-Za-z-]*))*
   )
)?
(?<BuildMetadataTagWithSeparator>
  \+(?<BuildMetadataTag>[\dA-Za-z-]+(\.[\dA-Za-z-]*)*)
)?$