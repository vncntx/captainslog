<#
.SYNOPSIS
Install dependencies

.DESCRIPTION
Ensure all dependencies and tools are installed

.EXAMPLE
Install-Project
#>
Function Install-Project {
	Confirm-Environment
	Write-Info "checking dependencies"
	If ((go mod verify) -And (Assert-ExitCode 0)) {
		Write-Success "all modules verified"
	}
	Else {
		Write-Warning "failed to verify modules"
	}
	Install-Hooks
	Write-Success "project installed"
}

<#
.SYNOPSIS
Verify the build environment

.DESCRIPTION
Verify the build environment is set up correctly

.EXAMPLE
Confirm-Environment
#>
Function Confirm-Environment {
	Write-Info "checking environment"
	if (-Not (Get-Command -Name go -ErrorAction SilentlyContinue)) {
		Write-Error "go is required"
	}
	if ($env:GO111MODULE -ne "on") {
		Write-Warning "go modules should be enabled"
	} else {
		Write-Success "go modules are enabled"
	}
	$GO_VERSION = "1.13"
	if (-Not (go version | Select-String -SimpleMatch "go$GO_VERSION")) {
		Write-Warning "go v$GO_VERSION should be installed"
	} else {
		Write-Success "go v$GO_VERSION is installed"
	}
}

<#
.SYNOPSIS
Install git hooks

.DESCRIPTION
Copy this project's git hooks into the .git directory

.EXAMPLE
Install-Hooks
#>
Function Install-Hooks {
	New-Item -Type Directory -Force (Join-Path ".git" "hooks") > $null
	ForEach ($file in (Get-ChildItem (Join-Path "hooks" "*.*"))) {
		# Get name without extension
		$name = ($file.Name -Split '\.')[0]
		$dest = (Join-Path ".git" "hooks" $name)
		Write-Info "installing $name hooks"

		Copy-Item $file $dest
		if (Get-Command chmod -ErrorAction SilentlyContinue) {
			chmod +x $dest
		}
	}
	Write-Success "git hooks installed"
}

<#
.SYNOPSIS
Format code

.DESCRIPTION
Format all Go code in this project

.EXAMPLE
Format-Project
#>
Function Format-Project {
	go fmt "./..."
	Write-Success "go style guide applied"
	go mod tidy
	Write-Success "dependencies tidied up"
}

<#
.SYNOPSIS
Run unit tests

.DESCRIPTION
Run unit tests for this project and all its packages

.EXAMPLE
Invoke-Tests
#>
Function Invoke-Tests {
	$module = (Get-GoModule)
	Write-Info "running tests"
	$output = (go test "./..." --cover)
	$failed = 0
	$total = 0
	ForEach ($line In ($output | Select-String -AllMatches -Pattern "\w\s+$module")) {
		$status, $module, $coverage = Get-TestDetails($line)
		$total++
		If ($status -Like "ok*") {
			Write-Pass "$module , $coverage" 	
		} Else {
			$failed++
			Write-Failure "$module , $coverage"
		}
	}
	Write-Output " ----".PadLeft(5)
	If ((Assert-ExitCode 0) -and ($failed -eq 0)) {
		If ($total -gt 0) {
			Write-Success "all tests passing"
		} Else {
			Write-Warning "no unit tests"
		}
	} Else {
		Write-Failure "$failed of $total packages failing"
		Write-Output ""
		Write-Output $output
	}
}

Function Get-TestDetails($testOutput) {
	$parts = ($line -Split "\t")
	Return $parts[0].ToUpper(), $parts[1], $parts[-1]
}

Function Get-GoModule {
	Return ((Get-Content go.mod)[0] -Split " ")[-1]
}

<#
.SYNOPSIS
Examine code for common flaws

.DESCRIPTION
Examine Go source code and report suspicious constructs

.EXAMPLE
Invoke-Checks
#>

Function Invoke-Checks {
	Write-Info "examining packages"
	go vet "./..."
	If (Assert-ExitCode 0) {
		Write-Success "no problems detected"
	} Else {
		Write-Failure "detected a few problems"
	}
}

<#
.SYNOPSIS
Run benchmarks

.DESCRIPTION
See how captainslog performs

.EXAMPLE
Invoke-Benchmarks
#>
Function Invoke-Benchmarks {
	Write-Info "running benchmark tests"
	$startTime = (Get-Date)
	$output = (go run ./benchmarks/)
	$executionTime = (Get-Date).Subtract($startTime).TotalMilliseconds / 60
	If (Assert-ExitCode 0) {
			Write-Success "benchmarks completed in $executionTime s"
			Write-Output $output[-5..-1]
	} Else {
			Write-Error "there was an error while running the benchmarks"
			Write-Output $output
	}
}


<#
.SYNOPSIS
Run captainslog demo

.DESCRIPTION
See captainslog in action

.EXAMPLE
Invoke-Demo
#>
Function Invoke-Demo {
	go run ./demo/
}

##########################################################################################

Export-ModuleMember -Function Format-Project
Export-ModuleMember -Function Install-Project
Export-ModuleMember -Function Invoke-Benchmarks
Export-ModuleMember -Function Invoke-Checks
Export-ModuleMember -Function Invoke-Demo
Export-ModuleMember -Function Invoke-Tests

##########################################################################################
#                                 Utility Functions                                      #
##########################################################################################

function Assert-ExitCode($expectedExitCode)
{
	return ($LASTEXITCODE -eq $expectedExitCode)
}

Function Write-Success($message) {
	Write-Host -NoNewline -ForegroundColor Green "ok".PadLeft(5)
	Write-Host " : $message"
}

Function Write-Pass($message) {
	Write-Host -NoNewline -ForegroundColor Green "pass".PadLeft(5)
	Write-Host " : $message"
}
Function Write-Failure($message) {
	Write-Host -NoNewline -ForegroundColor Red "fail".PadLeft(5)
	Write-Host " : $message"
}

Function Write-Warning($message) {
	Write-Host -NoNewline -ForegroundColor Yellow "warn".PadLeft(5)
	Write-Host " : $message"
}

Function Write-Error($message) {
	Write-Host -NoNewline -ForegroundColor Red "error".PadLeft(5)
	Write-Host " : $message"
	Throw $message
}

Function Write-Info($message) {
	Write-Host -NoNewline -ForegroundColor CYAN "info".PadLeft(5)
	Write-Host " : $message"
}