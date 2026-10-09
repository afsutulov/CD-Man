//go:build !windows && cgo

package platform

/*
#cgo linux LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

typedef struct { int x,y,w,h; } Rect;
typedef struct { int freq; uint16_t format; uint8_t channels,silence; uint16_t samples,padding; uint32_t size; void *callback,*userdata; } AudioSpec;
static void *library,*window,*renderer,*texture,*joystick;
static int tw,th,fullscreen=1;
static uint32_t audioDevice;
static int (*Init)(uint32_t);
static void (*Quit)(void);
static const char *(*GetError)(void);
static void *(*CreateWindow)(const char*,int,int,int,int,uint32_t);
static void (*DestroyWindow)(void*);
static int (*SetWindowFullscreen)(void*,uint32_t);
static void *(*CreateRenderer)(void*,int,uint32_t);
static void (*DestroyRenderer)(void*);
static void *(*CreateTexture)(void*,uint32_t,int,int,int);
static void (*DestroyTexture)(void*);
static int (*UpdateTexture)(void*,const Rect*,const void*,int);
static int (*SetRenderDrawColor)(void*,uint8_t,uint8_t,uint8_t,uint8_t);
static int (*RenderClear)(void*);
static int (*RenderCopy)(void*,void*,const Rect*,const Rect*);
static void (*RenderPresent)(void*);
static int (*GetRendererOutputSize)(void*,int*,int*);
static int (*PollEvent)(void*);
static int (*ShowCursor)(int);
static int (*SetHint)(const char*,const char*);
static uint32_t (*OpenAudioDevice)(const char*,int,const AudioSpec*,AudioSpec*,int);
static void (*CloseAudioDevice)(uint32_t);
static void (*PauseAudioDevice)(uint32_t,int);
static int (*QueueAudio)(uint32_t,const void*,uint32_t);
static uint32_t (*GetQueuedAudioSize)(uint32_t);
static void (*ClearQueuedAudio)(uint32_t);
static int (*NumJoysticks)(void);
static void *(*JoystickOpen)(int);
static void (*JoystickClose)(void*);
static int16_t (*JoystickGetAxis)(void*,int);
static uint8_t (*JoystickGetButton)(void*,int);

#define LOAD(name) do { *(void **)(&name)=dlsym(library,"SDL_" #name); if(!name)return "Missing SDL_" #name; } while(0)
static const char *startSDL(void) {
 const char *names[]={"libSDL2-2.0.so.0","libSDL2.so","libSDL2.dylib","/opt/homebrew/lib/libSDL2.dylib","/usr/local/lib/libSDL2.dylib","/Library/Frameworks/SDL2.framework/SDL2",NULL};
 for(int i=0;names[i]&&!library;i++)library=dlopen(names[i],RTLD_NOW|RTLD_LOCAL);
 if(!library)return "SDL2 runtime library is required. See README.md.";
 LOAD(Init);LOAD(Quit);LOAD(GetError);LOAD(CreateWindow);LOAD(DestroyWindow);LOAD(SetWindowFullscreen);LOAD(CreateRenderer);LOAD(DestroyRenderer);LOAD(CreateTexture);LOAD(DestroyTexture);LOAD(UpdateTexture);LOAD(SetRenderDrawColor);LOAD(RenderClear);LOAD(RenderCopy);LOAD(RenderPresent);LOAD(GetRendererOutputSize);LOAD(PollEvent);LOAD(ShowCursor);LOAD(SetHint);LOAD(OpenAudioDevice);LOAD(CloseAudioDevice);LOAD(PauseAudioDevice);LOAD(QueueAudio);LOAD(GetQueuedAudioSize);LOAD(ClearQueuedAudio);LOAD(NumJoysticks);LOAD(JoystickOpen);LOAD(JoystickClose);LOAD(JoystickGetAxis);LOAD(JoystickGetButton);
 // SDL_INIT_AUDIO=0x10, VIDEO=0x20, JOYSTICK=0x200.
 if(Init(0x10|0x20|0x200)<0)return GetError();
 SetHint("SDL_RENDER_SCALE_QUALITY","0");
 window=CreateWindow("CD-Man 2.0",0x2fff0000,0x2fff0000,960,720,0x1001|0x20|0x2000);
 if(!window)return GetError();
 renderer=CreateRenderer(window,-1,0);
 if(!renderer)return GetError();
 ShowCursor(0);
 AudioSpec want={0},got={0};want.freq=48000;want.format=0x8010;want.channels=1;want.samples=1024;
 audioDevice=OpenAudioDevice(NULL,0,&want,&got,0);if(!audioDevice)return GetError();PauseAudioDevice(audioDevice,0);
 if(NumJoysticks()>0)joystick=JoystickOpen(0);
 return NULL;
}
static void stopSDL(void){if(joystick&&JoystickClose)JoystickClose(joystick);if(audioDevice&&CloseAudioDevice)CloseAudioDevice(audioDevice);if(texture&&DestroyTexture)DestroyTexture(texture);if(renderer&&DestroyRenderer)DestroyRenderer(renderer);if(window&&DestroyWindow)DestroyWindow(window);if(Quit)Quit();if(library)dlclose(library);}
static int pollSDL(int *scan,int *sym,int *mod){union {uint64_t align;uint8_t b[64];} e;while(PollEvent(&e)){uint32_t kind;memcpy(&kind,e.b,4);if(kind==0x100)return -1;if(kind==0x300){memcpy(scan,e.b+16,4);memcpy(sym,e.b+20,4);uint16_t v;memcpy(&v,e.b+24,2);*mod=v;if(*scan==40&&(*mod&0x300)){fullscreen=!fullscreen;SetWindowFullscreen(window,fullscreen?0x1001:0);continue;}return 1;}}return 0;}
static const char *drawSDL(const void *pixels,int w,int h){if(w!=tw||h!=th){if(texture)DestroyTexture(texture);texture=CreateTexture(renderer,0x16762004,1,w,h);tw=w;th=h;if(!texture)return GetError();}if(UpdateTexture(texture,NULL,pixels,w*4)<0)return GetError();int cw,ch;GetRendererOutputSize(renderer,&cw,&ch);int dw=cw,dh=cw*3/4;if(dh>ch){dh=ch;dw=ch*4/3;}Rect dst={(cw-dw)/2,(ch-dh)/2,dw,dh};SetRenderDrawColor(renderer,0,0,0,255);RenderClear(renderer);if(RenderCopy(renderer,texture,NULL,&dst)<0)return GetError();RenderPresent(renderer);return NULL;}
static const char *soundSDL(const void *samples,uint32_t bytes){if(GetQueuedAudioSize(audioDevice)>48000/2)ClearQueuedAudio(audioDevice);if(QueueAudio(audioDevice,samples,bytes)<0)return GetError();return NULL;}
static int joySDL(uint16_t *x,uint16_t *y,uint8_t *buttons){if(!joystick)return 0;*x=(uint16_t)((int)JoystickGetAxis(joystick,0)+32768);*y=(uint16_t)((int)JoystickGetAxis(joystick,1)+32768);*buttons=JoystickGetButton(joystick,0)|(JoystickGetButton(joystick,1)<<1);return 1;}
*/
import "C"

import (
	"cdman2/internal/game"
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func ShowError(s string) { fmt.Println(s) }

// Cocoa requires the native window loop on the process's initial thread.
func init() { runtime.LockOSThread() }

var usbToDOS = map[int]byte{
	4: 0x1e, 5: 0x30, 6: 0x2e, 7: 0x20, 8: 0x12, 9: 0x21, 10: 0x22, 11: 0x23, 12: 0x17, 13: 0x24, 14: 0x25, 15: 0x26, 16: 0x32, 17: 0x31, 18: 0x18, 19: 0x19, 20: 0x10, 21: 0x13, 22: 0x1f, 23: 0x14, 24: 0x16, 25: 0x2f, 26: 0x11, 27: 0x2d, 28: 0x15, 29: 0x2c,
	30: 2, 31: 3, 32: 4, 33: 5, 34: 6, 35: 7, 36: 8, 37: 9, 38: 10, 39: 11, 40: 0x1c, 41: 1, 42: 0xe, 43: 0xf, 44: 0x39, 45: 0xc, 46: 0xd, 47: 0x1a, 48: 0x1b, 49: 0x2b, 51: 0x27, 52: 0x28, 53: 0x29, 54: 0x33, 55: 0x34, 56: 0x35,
	58: 0x3b, 59: 0x3c, 60: 0x3d, 61: 0x3e, 62: 0x3f, 63: 0x40, 64: 0x41, 65: 0x42, 66: 0x43, 67: 0x44, 68: 0x57, 69: 0x58,
	73: 0x52, 74: 0x47, 75: 0x49, 76: 0x53, 77: 0x4f, 78: 0x51, 79: 0x4d, 80: 0x4b, 81: 0x50, 82: 0x48, 88: 0x1c, 89: 0x4f, 90: 0x50, 91: 0x51, 92: 0x4b, 93: 0x4c, 94: 0x4d, 95: 0x47, 96: 0x48, 97: 0x49, 98: 0x52, 99: 0x53,
}

func Run(m *game.Machine, save func() error) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	defer C.stopSDL()
	if err := C.startSDL(); err != nil {
		return fmt.Errorf("%s", C.GoString(err))
	}
	deadline := time.Now()
	for !m.Halted() {
		for {
			var scan, sym, mod C.int
			state := C.pollSDL(&scan, &sym, &mod)
			if state < 0 {
				return nil
			}
			if state == 0 {
				break
			}
			code, ok := usbToDOS[int(scan)]
			if !ok {
				continue
			}
			ascii := byte(0)
			if sym > 0 && sym < 128 {
				ascii = byte(sym)
			}
			if ascii >= 'a' && ascii <= 'z' {
				if int(mod)&0xc0 != 0 {
					ascii -= 96
				} else if (int(mod)&3 != 0) != (int(mod)&0x2000 != 0) {
					ascii -= 32
				}
			}
			if int(mod)&3 != 0 {
				plain := "1234567890-=[]\\;',./`"
				shifted := "!@#$%^&*()_+{}|:\"<>?~"
				for i := range plain {
					if ascii == plain[i] {
						ascii = shifted[i]
						break
					}
				}
			}
			m.Key(code, ascii)
		}
		var x, y C.uint16_t
		var buttons C.uint8_t
		present := C.joySDL(&x, &y, &buttons)
		m.Joystick(present != 0, uint16(x), uint16(y), byte(buttons))
		m.RunCycles(game.ClockHz / 70)
		samples := m.Audio()
		if len(samples) > 0 {
			if e := C.soundSDL(unsafe.Pointer(&samples[0]), C.uint32_t(len(samples)*2)); e != nil {
				return fmt.Errorf("%s", C.GoString(e))
			}
		}
		if err := save(); err != nil {
			return err
		}
		im := m.Frame()
		if e := C.drawSDL(unsafe.Pointer(&im.Pix[0]), C.int(im.Bounds().Dx()), C.int(im.Bounds().Dy())); e != nil {
			return fmt.Errorf("%s", C.GoString(e))
		}
		deadline = deadline.Add(time.Second / 70)
		if d := time.Until(deadline); d > 0 {
			time.Sleep(d)
		} else if d < -time.Second/4 {
			deadline = time.Now()
		}
	}
	return m.Error()
}
