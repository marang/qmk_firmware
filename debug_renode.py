#!/usr/bin/env python3

import subprocess
import sys
import time

def test_renode_basic():
    """Test basic Renode functionality"""
    
    # Try a simple Renode command to see if it works at all
    try:
        result = subprocess.run([
            'renode', '--disable-xwt', 
            '-e', 'help'
        ], capture_output=True, text=True, timeout=10)
        
        print("Renode basic command output:")
        print(result.stdout[:500] if result.stdout else "No stdout")
        if result.stderr:
            print("Renode stderr:", result.stderr[:500])
            
        return result.returncode == 0
        
    except Exception as e:
        print(f"Error running Renode basic test: {e}")
        return False

def test_firmware_loading():
    """Test if firmware can be loaded in Renode"""
    
    # Create a minimal renode script to test just firmware loading
    script_content = '''
mach create "test-machine"
sysbus LoadELF @build/firmware.elf
cpu SingleStep
quit
'''
    
    script_file = '/tmp/test_renode_script.repl'
    with open(script_file, 'w') as f:
        f.write(script_content)
    
    try:
        result = subprocess.run([
            'renode', '--disable-xwt',
            '-e', f'machine LoadPlatformDescription @{script_file}'
        ], capture_output=True, text=True, timeout=15)
        
        print("Firmware loading test:")
        print("Return code:", result.returncode)
        print("STDOUT:", result.stdout[-500:] if result.stdout else "None")
        if result.stderr:
            print("STDERR:", result.stderr[-500:] if result.stderr else "None")
            
        return result.returncode == 0
        
    except Exception as e:
        print(f"Error in firmware loading test: {e}")
        return False

if __name__ == "__main__":
    print("=== Renode Debug Tests ===")
    
    print("\n1. Testing basic Renode functionality...")
    basic_ok = test_renode_basic()
    print(f"Basic test: {'PASS' if basic_ok else 'FAIL'}")
    
    print("\n2. Testing firmware loading...")
    firmware_ok = test_firmware_loading()
    print(f"Firmware test: {'PASS' if firmware_ok else 'FAIL'}")
    
    print(f"\n=== Summary ===")
    print(f"Overall: {'PASS' if basic_ok and firmware_ok else 'FAIL'}")