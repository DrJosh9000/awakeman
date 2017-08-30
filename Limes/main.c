#include <Quickdraw.h>
#include <MacMemory.h>
#include <Sound.h>
#include <Events.h>
#include <Fonts.h>
#include <NumberFormatting.h>
#include <Resources.h>
#include <ToolUtils.h>
#include <OSUtils.h>

#include <math.h>
#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>


// --------------------------- Consts

// I don't know where these are defined. Inside Macintosh lists different codes
// for each of the keys. 
// In the end was easier to knock up a quick program for dumping keyboard state
// each second.
#define kQKeyMap 11
#define	kIKeyMap 37
#define	kLKeyMap 34
#define	kJKeyMap 33
#define	kKKeyMap 47

const short playerWidth = 16;
const short playerHeight = 16;


// --------------------------- Types

typedef struct {
	Rect src, dst;
	// Rect maskSrc;
	short x, y;
	short dx, dy;
	short frame, frameStart, animask;
} player;

struct lime;

struct lime {
	Rect dst;
	int expiry;
	struct lime *next;
};
typedef struct lime lime;


// --------------------------- Vars

WindowPtr win;
Rect offPartsRect;
GrafPtr offPartsPtr;
BitMap offPartsBits;

Rect theBounds;
player thePlayer;
lime *theLimes;  // Linked list with dummy first node.
int numLimes;
int theScore;
Rect limeSrcRect;


// --------------------------- Player methods

// Recompute the src and dst rects for the player object.
void compute_player_rects() {
	short fx = thePlayer.frame*playerWidth;
	SetRect(&thePlayer.src, fx, 0, fx+playerWidth, playerHeight);
	//SetRect(&thePlayer.maskSrc, fx, 16, fx+playerWidth, playerHeight);
	SetRect(&thePlayer.dst, thePlayer.x, thePlayer.y, thePlayer.x+playerWidth, thePlayer.y+playerHeight);
}

// Un-draw the player.
void clear_player() {
	// srcBic = "bit clear", clearing bits that are set in the src. Useful for clearing.
	CopyBits(&offPartsBits, &win->portBits, &thePlayer.src, &thePlayer.dst, srcBic, nil);
}

// Draw the player.
void draw_player() {
	// srcCopy just copies.
	CopyBits(&offPartsBits, &win->portBits, &thePlayer.src, &thePlayer.dst, srcCopy, nil);
}

// Clear the player, recompute the rects, and then draw the player.
void redraw_player() {
	// TODO: should reimplement - draw the delta into a small 
	// backbuffer then copying to the screen in 1 call.
	clear_player();
	compute_player_rects();
	draw_player();
}

// Update the player state by 1 tick.
// Returns true if the player changed in some (visible) way.
bool update_player() {
	if (thePlayer.dx != 0 || thePlayer.dy != 0) { // Draw the player walking.
		thePlayer.x += thePlayer.dx;
		thePlayer.y += thePlayer.dy;
		// Constrain to theBounds
		if (thePlayer.x < theBounds.left) {
			thePlayer.x = theBounds.left;
		}
		if (thePlayer.x > theBounds.right-playerWidth) {
			thePlayer.x = theBounds.right-playerWidth;
		}
		if (thePlayer.y < theBounds.top) {
			thePlayer.y = theBounds.top;
		}
		if (thePlayer.y > theBounds.bottom-playerHeight) {
			thePlayer.y = theBounds.bottom-playerHeight;
		}
		// One animation frame per 4 ticks.
		// A power-of-2 number of frames per animation.
		thePlayer.frame = ((TickCount() >> 2) & thePlayer.animask) + thePlayer.frameStart;
		return true;
	} 
	if (thePlayer.frame != 0) { // Need to redraw the player, standing.
		thePlayer.frame = 0;
		return true;
	}
	return false;
}


// --------------------------- Lime methods

// Draw all the limes.
void draw_limes() {
	lime *l = theLimes->next;
	while (l != nil) {
		CopyBits(&offPartsBits, &win->portBits, &limeSrcRect, &l->dst, srcCopy, nil);
		l = l->next;
	}
}

// Un-draw a lime.
void clear_lime(lime *l) {
	CopyBits(&offPartsBits, &win->portBits, &limeSrcRect, &l->dst, srcBic, nil);	
}

// Adds a lime and returns a pointer to it.
lime* add_lime() {
	numLimes++;
	lime *l = theLimes->next;
	theLimes->next = (lime *)calloc(1, sizeof(lime));
	theLimes->next->next = l;
	return theLimes->next;
}

// Remove the lime *after* the given lime.
void remove_lime(lime *prev) {
	lime *l = prev->next;
	if (l == nil) {
		return;
	}
	numLimes--;	
	prev->next = l->next;
	clear_lime(l);
	free(l);
}

// Returns lime in the list *before* the intersecting lime.
lime* player_intersect_limes() {
	lime *l = theLimes;
	while (l->next != nil) {
		Rect ignore;
		if (SectRect(&l->next->dst, &thePlayer.dst, &ignore)) {
			return l;
		}
		l = l->next;
	}
}

void generate_limes(int n) {
	for (short i=0; i<n; i++) {
		lime *l = add_lime();
		l->expiry = TickCount() + (rand() % (10*60)) + 120;		
		short x = (rand() % (theBounds.right-theBounds.left-16))+theBounds.left;
		short y = (rand() % (theBounds.bottom-theBounds.top-16))+theBounds.top;
		SetRect(&l->dst, x, y, x+16, y+16);
	}
}

int expire_limes() {
	// Remove any expired limes.
	int expired = 0;
	lime *l = theLimes;
	while (l->next != nil) {
		if (TickCount() > l->next->expiry) {
			remove_lime(l);
			expired++;
			continue;
		}
		l = l->next;
	}
	return expired;
}

// --------------------------- Input methods.

// Returns true if the game keeps going.
bool read_input() {
	KeyMap theKeyMap;
	GetKeys(theKeyMap);
	if (BitTst(&theKeyMap, kQKeyMap)) {
		// Quit
		return false;
	}
	thePlayer.dx = 0;
	thePlayer.dy = 0;
	if (BitTst(&theKeyMap, kJKeyMap)) {
		// Left
		thePlayer.frameStart = 5;
		thePlayer.animask = 0x03;
		thePlayer.dx = -1;
	}
	if (BitTst(&theKeyMap, kLKeyMap)) {
		// Right
		thePlayer.frameStart = 1;
		thePlayer.animask = 0x03;
		thePlayer.dx = 1;
	}
	if (BitTst(&theKeyMap, kIKeyMap)) {
		// Up
		thePlayer.frameStart = 9;
		thePlayer.animask = 0x07;
		thePlayer.dy = -1;
	}
	if (BitTst(&theKeyMap, kKKeyMap)) {
		// Down
		thePlayer.frameStart = 9;	
		thePlayer.animask = 0x07;		
		thePlayer.dy = 1;
	}
	return true;
}

// --------------------------- Other bits

void draw_score() {
	MoveTo(4, 14);
	char tmp[22];
	sprintf(tmp, "\pScore: %d", theScore);
	Rect r;
	SetRect(&r, 4, 4, 4+StringWidth(tmp), 14);
	EraseRect(&r);
	DrawString(tmp);
}

void draw_instructions() {
	MoveTo(4, theBounds.bottom+12);
	DrawString("\pq - quit; i,j,k,l - up, left, down, right");
}


// --------------------------- The program

int main() {
	// Standard random stuff
	srand(TickCount());

	// Initialise (most of) ALL THE THINGS
	theScore = 0;

	thePlayer.x = 60;
	thePlayer.y = 40;
	thePlayer.frame = 0;
	thePlayer.dx = 0;
	thePlayer.dy = 0;
	compute_player_rects();

	SetRect(&limeSrcRect, 0, 32, 16, 48);

	theLimes = (lime*)calloc(1, sizeof(lime));	

    InitGraf(&qd.thePort);
    InitFonts();
    InitWindows();
    InitMenus();
	FlushEvents(everyEvent, 0);

	// Make a window - make it slightly smaller than the screen below the menu bar.
	Rect r = qd.screenBits.bounds;
    SetRect(&r, r.left + 5, r.top + 45, r.right - 5, r.bottom - 5);
	win = NewWindow(NULL, &r, "\pAwakeman Picks Up Some Limes", true, 0, (WindowPtr)-1, false, 0);
	SetPort(win);
	r = win->portRect;
	EraseRect(&r);
	
	SetRect(&theBounds, r.left, r.top+16, r.right, r.bottom-16);
	//FillRect(&theBounds, &win->fillPat);

	// Create a Graf for storing the sprite atlas, and make it the 
	// same size as the PICT we're going to load from resources.
	offPartsPtr = (GrafPtr)(NewPtr(sizeof(GrafPort)));
	OpenPort(offPartsPtr);
	const short mapWidth = 426;
	const short mapHeight = 128;
	
	SetRect(&offPartsRect, 0, 0, mapWidth, mapHeight);
	short partsRowBytes = ((offPartsRect.right - offPartsRect.left + 15) / 16) * 2;
	
	offPartsBits.rowBytes = partsRowBytes;
	offPartsBits.bounds = offPartsRect;
	offPartsBits.baseAddr = NewPtr((long)offPartsBits.rowBytes * 
		(offPartsRect.bottom - offPartsRect.top));
	SetPortBits(&offPartsBits);
	ClipRect(&offPartsRect);
	EraseRect(&offPartsRect);

	// GetResource: searches all the resource maps in memory for the specified resource.
	// Get1Resource: searches the current resource file only. (Not in original Inside Macintosh.)
	//PicHandle hSprites = (PicHandle)Get1Resource('PICT', 128);  
	// (Single-quote "string" -> unsigned 32-bit int literal? Yep!)
	// But wait! Quickdraw has a handy function for doing this:
	PicHandle hSprites = GetPicture(128);
	
	// Draw the picture into the offParts...
	DrawPicture(hSprites, &offPartsRect);

	// Relinquish the PICT back unto the void.
	ReleaseResource((Handle)hSprites);

	// Now ready for user input - change cursor to a regular pointer to
	// indicate we're done.
	InitCursor();

	// Go back to drawing on the window.
	SetPort(win);

	// Initial view.
	generate_limes(5);
	draw_instructions();
	draw_player();
	draw_limes();
	draw_score();

	// Main loop
	int waitUntil = TickCount() + 1;
	while(read_input()) {
		while (TickCount() < waitUntil) 
			;
		waitUntil = TickCount() + 1;
		
		int gen = expire_limes();
		if (numLimes == 0) {
			gen = 5;
		}
		if (gen > 0) {
			generate_limes(gen);
			draw_limes();
		}
		
		// Player movement?
		if (update_player()) {
			// Intersecting any limes?
			lime *l = player_intersect_limes();
			if (l != nil) {
				theScore++;				
				remove_lime(l);
				draw_score();
			}
			redraw_player();
		}
	}
	FlushEvents(everyEvent, -1);
	return 0;
}